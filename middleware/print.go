package middleware

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Print() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := time.Now()

		header, _ := json.Marshal(c.Request.Header)
		fmt.Println("-----------------request print start----------------------")
		fmt.Println("Print ClientIP: ", c.ClientIP())
		fmt.Println("Print Request Header:", string(header))
		fmt.Println("Print Response Status:", c.Writer.Status())
		fmt.Println("Print Request Host:", c.Request.Host)
		fmt.Println("Print Request URI:", c.Request.RequestURI)
		fmt.Println("Print Request Method: ", c.Request.Method)
		fmt.Println("Print Request Query:", c.Request.URL.Query())
		fmt.Println("Print Request Body:", c.Request.MultipartForm)
		fmt.Println("-----------------request print end------------------------")
		c.Next()
		e := time.Now()

		latency := e.Sub(s)
		fmt.Printf("Print Response Time: %13v\n", latency)
	}
}
