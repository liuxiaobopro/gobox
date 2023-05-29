package excel

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRead_Read(t *testing.T) {
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

	excelRead := NewRead(
		WithFile(file.Filename),
		WithLu(Cell{Col: "A", Row: 2}),
		WithRd(Cell{Col: "F", Row: 2}),
	)

	var (
		res [][]string
		err error
	)

	if res, err = excelRead.Read(); err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"msg": res,
	})
}
