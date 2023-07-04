package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	ginx "github.com/liuxiaobopro/gobox/gin"
	logx "github.com/liuxiaobopro/gobox/log"
)

const MAX_PRINT_BODY_LEN = 512000

type bodyLogWriter struct {
	gin.ResponseWriter
	bodyBuf *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	// 内存拷贝
	w.bodyBuf.Write(b)
	return w.ResponseWriter.Write(b)
}

func Print(logger *logx.Gin) gin.HandlerFunc {
	return func(c *gin.Context) {
		s := time.Now()

		var reqBody []byte

		// 获取Content-Type
		contentType := c.ContentType()
		if contentType == "application/json" {
			// // 获取请求体
			// reqBody, err := ioutil.ReadAll(c.Request.Body)
			// // var req map[string]interface{}
			// // err := json.NewDecoder(c.Request.Body).Decode(&req)
			// if err != nil {
			// 	// 处理或记录错误
			// 	c.AbortWithStatus(500)
			// 	return
			// }
			// // reqBody, _ := json.Marshal(req)

			// // 重置请求体
			// c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))
			s := ginx.GetBody(c)
			reqBody = []byte(s)
		}

		// 记录请求日志
		header, _ := json.Marshal(c.Request.Header)
		param := ginx.GetParam(c)
		str := fmt.Sprintf(`
    Request ClientIP: %s
    Request Header: %s
    Request Host: %s
    Request URI: %s
    Request Method: %s
    Request Param: %v
    Request Body: %s
`,
			c.ClientIP(),
			string(header),
			c.Request.Host,
			c.Request.RequestURI,
			c.Request.Method,
			param,
			string(reqBody),
		)
		logger.Infof(c, str)

		// 捕获响应数据
		strBody := ""
		blw := bodyLogWriter{bodyBuf: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		// 继续处理请求
		c.Next()

		// 获取响应数据
		strBody = strings.Trim(blw.bodyBuf.String(), "\n")
		if len(strBody) > MAX_PRINT_BODY_LEN {
			// strBody = strBody[:(MAX_PRINT_BODY_LEN - 1)]
			strBody = "body too long, not print"
		}

		// 记录响应日志
		str1 := fmt.Sprintf(`
    Response Status: %d
    Response Data : %s
`,
			c.Writer.Status(),
			strBody,
		)
		logger.Infof(c, str1)

		e := time.Now()
		latency := e.Sub(s)
		logger.Infof(c, "Print latency: %s", latency)
	}
}
