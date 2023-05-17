package file

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestNewQiniu(t *testing.T) {
	r := gin.Default()
	r.POST("/upload", func(c *gin.Context) {
		file, fileHeader, err := c.Request.FormFile("file")
		if err != nil {
			t.Errorf("c.Request.FormFile(\"file\") error(%v)", err)
			c.JSON(200, gin.H{
				"code": 1,
				"msg":  "c.Request.FormFile(\"file\") error",
			})
			return
		}

		qiniu := NewQiniu(
			WithKey("7cQPM5OR753SZYtNXXWgjt_0bVsiI8_mIj2ng90g", "rMlzsj2Wq9WZDUjB60J9Lffvjy8I7Q7ngxv3GZf8"),
			WithBucket("liuxiaobo-img"),
			WithImgUrl("img.liuxiaobo.net.cn"),
			WithFilePath("./img/"),
		)

		// func (f *File) UploadFile(file multipart.File, fileHeader *multipart.FileHeader) (string, error)
		imgUrl, err := qiniu.UploadFile(file, fileHeader)

		if err != nil {
			t.Errorf("qiniu.UploadFile(file, fileHeader) error(%v)", err)
			c.JSON(200, gin.H{
				"code": 1,
				"msg":  fmt.Sprintf("qiniu.UploadFile(file, fileHeader) error(%s)", err.Error()),
			})
			return
		}

		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "success",
			"data": gin.H{
				"imgUrl": imgUrl,
			},
		})
	})

	r.Run(":8080")
}
