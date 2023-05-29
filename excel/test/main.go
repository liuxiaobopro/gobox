package main

import (
	"github.com/gin-gonic/gin"

	excelx "github.com/liuxiaobopro/gobox/excel"
)

func main() {
	r := gin.Default()
	r.POST("/upload", upload)
	r.Run(":8081")
}

// upload 上传
func upload(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(200, gin.H{
				"msg": err.(error).Error(),
			})
		}
	}()

	file, _ := c.FormFile("file")
	if err := c.SaveUploadedFile(file, file.Filename); err != nil {
		panic(err)
	}

	excelRead := excelx.NewRead(
		excelx.WithFile(file.Filename),
		excelx.WithLu(excelx.Cell{Col: "A", Row: 2}),
		excelx.WithRd(excelx.Cell{Col: "F"}),
	)

	var (
		res [][]string
		err error
	)

	if res, err = excelRead.Read(); err != nil {
		panic(err)
	}

	excelRead.Print()

	c.JSON(200, gin.H{
		"msg": res,
	})
}
