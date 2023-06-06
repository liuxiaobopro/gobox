package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	ossx "github.com/liuxiaobopro/gobox/alibaba/oss"
)

func main() {
	r := gin.Default()

	r.POST("/upload", func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()

		f, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// 保存在指定路径
		if err := c.SaveUploadedFile(f, f.Filename); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		oss := ossx.NewOss(
			ossx.WithAssessKey("xxx", "xxx"),
			ossx.WithEndpoint("https://oss-cn-beijing.aliyuncs.com"),
			ossx.WithBucket("xxx"),
		)

		if err := oss.CreateClient(); err != nil {
			panic(err)
		}

		src, err := f.Open()
		if err != nil {
			panic(err)
		}
		defer src.Close()

		path := "aliupload/" + f.Filename

		if url, err := oss.PutObject(path, src); err != nil {
			panic(err)
		} else {
			c.String(200, url)
		}
	})

	_ = r.Run(":8080")

}
