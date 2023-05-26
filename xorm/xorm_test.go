package xorm

import (
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/core"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

func TestNewGenXormDao(t *testing.T) {
	engine, err := xorm.NewEngine("mysql", "root:111111@tcp(192.168.31.75:3306)/today_earth?charset=utf8mb4")
	if err != nil {
		panic(err)
	}

	// f, err := os.Create("runtime/log/sql.log")
	if err != nil {
		panic(err)
	}
	// engine.SetLogger(log.NewSimpleLogger(f)) // 设置日志输出位置
	engine.ShowSQL(true)                                                     // 在控制台打印sql
	engine.SetLogLevel(log.LOG_DEBUG)                                        // 设置日志级别
	engine.SetMaxIdleConns(10)                                               // 设置连接池的空闲数大小
	engine.SetMaxOpenConns(100)                                              // 设置最大打开连接数
	engine.SetConnMaxLifetime(10)                                            // 设置连接的最大生命周期
	engine.SetTZLocation(time.Local)                                         // 设置时区
	engine.SetTableMapper(core.NewPrefixMapper(core.SnakeMapper{}, "qsgo_")) // 设置前缀

	gxd := NewGenXormDao(
		WithMysql(engine),
		WithDaoMysqlPath("./dao/mysql"),
		WithProject("demo"),
		WithPrefix("qsgo_"),
		WithProgramTemplatePath("./tpl/dao_default.tpl"),
		WithDefaultTemplatePath("./tpl/dao_program.tpl"),
	)
	if err := gxd.Gen(); err != nil {
		t.Error(err)
	}
}
