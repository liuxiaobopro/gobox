package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	replyx "github.com/liuxiaobopro/gobox/reply"
)

type IPRequestCounter struct {
	Counters map[string]int
	Lock     sync.Mutex
}

func IPRateLimit(counter *IPRequestCounter, maxRequests int, duration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		// 加锁，确保并发安全
		counter.Lock.Lock()
		defer counter.Lock.Unlock()

		// 检查IP请求计数器中是否存在当前IP
		if counter.Counters[ip] >= maxRequests {
			c.JSON(http.StatusTooManyRequests, replyx.TooManyReqErrT)
			c.Abort()
			return
		}

		// 增加当前IP的请求计数
		counter.Counters[ip]++

		// 使用定时器重置计数器
		time.AfterFunc(duration, func() {
			counter.Lock.Lock()
			defer counter.Lock.Unlock()

			counter.Counters[ip]--
		})

		// 继续处理请求
		c.Next()
	}
}
