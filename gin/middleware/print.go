package middleware

// 参考链接: https://zhuanlan.zhihu.com/p/94309327

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	logx "github.com/liuxiaobopro/gobox/log"
)

const MAX_PRINT_BODY_LEN = 512000

type bodyLogWriter struct {
	gin.ResponseWriter
	bodyBuf *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	//memory copy here!
	w.bodyBuf.Write(b)
	return w.ResponseWriter.Write(b)
}

func Print(logger *logx.Gin) gin.HandlerFunc {
	return func(c *gin.Context) {
		s := time.Now()
		reqBody, _ := c.GetRawData()
		header, _ := json.Marshal(c.Request.Header)
		str := `
    ClientIP: ` + c.ClientIP() + `
    Request Header: ` + string(header) + `
    Request Host: ` + c.Request.Host + `
    Request URI: ` + c.Request.RequestURI + `
    Request Method: ` + c.Request.Method + `
    Request Query: ` + fmt.Sprintf("%v", c.Request.URL.Query()) + `
    Request Body: ` + string(reqBody) + `
`
		logger.Infof(c, str)

		strBody := ""
		blw := bodyLogWriter{bodyBuf: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		c.Next()

		strBody = strings.Trim(blw.bodyBuf.String(), "\n")
		if len(strBody) > MAX_PRINT_BODY_LEN {
			strBody = strBody[:(MAX_PRINT_BODY_LEN - 1)]
		}

		str1 := `
    Response Status: ` + fmt.Sprintf("%d", c.Writer.Status()) + `
    Response Data : ` + strBody + `
        `
		logger.Infof(c, str1)

		e := time.Now()
		latency := e.Sub(s)
		logger.Infof(c, "Print latency: %s", latency)
	}
}
