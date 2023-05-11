package middleware

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	logx "github.com/liuxiaobopro/gobox/log"
)

func Print(logger *logx.Gin) gin.HandlerFunc {
	return func(c *gin.Context) {
		s := time.Now()
		header, _ := json.Marshal(c.Request.Header)
		str := `
		-----------------request print start----------------------
		Print ClientIP: ` + c.ClientIP() + `
		Print Request Header: ` + string(header) + `
		Print Response Status: ` + fmt.Sprintf("%d", c.Writer.Status()) + `
		Print Request Host: ` + c.Request.Host + `
		Print Request URI: ` + c.Request.RequestURI + `
		Print Request Method: ` + c.Request.Method + `
		Print Request Query: ` + fmt.Sprintf("%v", c.Request.URL.Query()) + `
		Print Request Body: ` + fmt.Sprintf("%v", c.Request.MultipartForm) + `
		-----------------request print end------------------------
		`
		logger.Infof(c, str)

		c.Next()

		e := time.Now()
		latency := e.Sub(s)
		logger.Infof(c, "Print latency: %s", latency)
	}
}
