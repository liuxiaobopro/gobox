package log

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/robfig/cron/v3"
)

type cut struct {
	mode       int    // 日志切割模式 (1: 按时间 2: 按大小)
	filepath   string // 日志路径 (例如: ./log/output.log)
	fileFolder string // 日志文件夹路径 (例如: ./log/)

	// 按时间

	cron string // 计划任务表达式 (例如: 0 0 0 * * * 每天0点0分0秒切割)

	// 按大小

	size int64 // 日志大小 (例如: 1024 * 1024 * 1024 1G)
}

var (
	ErrInvalidMode = errors.New("invalid mode")
)

type cutOption func(*cut)

func WithCron(cron string) cutOption {
	return func(c *cut) {
		c.cron = cron
	}
}

func WithSize(size int64) cutOption {
	return func(c *cut) {
		c.size = size
	}
}

func NewCut(mode int, filepath string, options ...cutOption) *cut {
	c := &cut{
		mode:     mode,
		filepath: filepath,
	}

	for _, option := range options {
		option(c)
	}

	c.fileFolder = filepath[:strings.LastIndex(filepath, "/")+1]

	if c.mode == 0 && c.cron == "" {
		c.mode = 1
		c.cron = "0 0 0 * * *"
	}

	return c
}

func (c *cut) Start() {
	// 判断文件是否存在
	if _, err := os.Stat(c.filepath); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "log file not exist: %s\n", c.filepath)
		return
	}

	switch c.mode {
	case 1:
		c.start1()
	case 2:
		c.start2()
	}
}

// start1 按时间切割
func (c *cut) start1() {
	c1 := cron.New()

	if _, err := c1.AddFunc(c.cron, func() {
		// 判断当前系统
		switch runtime.GOOS {
		case "linux":
			// 备份
			cmd := exec.Command("cp", c.filepath, fmt.Sprintf("%s_%s.log", c.filepath[:strings.LastIndex(c.filepath, ".")], time.Now().Format("20060102150405")))
			if _, err := cmd.Output(); err != nil {
				fmt.Fprintf(os.Stderr, "backup log file failed: %s\n", err)
				return
			}

			// 清空
			cmd = exec.Command(">", c.filepath)
			if _, err := cmd.Output(); err != nil {
				fmt.Fprintf(os.Stderr, "clear log file failed: %s\n", err)
				return
			}

		default:
			fmt.Fprintf(os.Stderr, "unsupported os: %s\n", runtime.GOOS)
		}
	}); err != nil {
		return
	}

	c1.Start()
	defer c1.Stop()

	select {}
}

// start2 按大小切割
func (c *cut) start2() {
	c1 := cron.New()

	if _, err := c1.AddFunc("*/1 * * * *", func() {
		// 判断当前系统
		switch runtime.GOOS {
		case "linux":
			// 获取文件大小
			fi, err := os.Stat(c.filepath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "get log file size failed: %s\n", err)
				return
			}

			// 判断文件大小
			if fi.Size() < c.size {
				return
			}

			// 备份
			cmd := exec.Command("cp", c.filepath, fmt.Sprintf("%s_%s.log", c.filepath[:strings.LastIndex(c.filepath, ".")], time.Now().Format("20060102150405")))
			if _, err := cmd.Output(); err != nil {
				fmt.Fprintf(os.Stderr, "backup log file failed: %s\n", err)
				return
			}

			// 清空
			cmd = exec.Command(">", c.filepath)
			if _, err := cmd.Output(); err != nil {
				fmt.Fprintf(os.Stderr, "clear log file failed: %s\n", err)
				return
			}

		default:
			fmt.Fprintf(os.Stderr, "unsupported os: %s\n", runtime.GOOS)
		}
	}); err != nil {
		return
	}

	c1.Start()
	defer c1.Stop()

	select {}
}
