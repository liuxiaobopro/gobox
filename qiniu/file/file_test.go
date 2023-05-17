package file

import (
	"fmt"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/storage"
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
			WithKey("xxx", "xxx"),
			WithBucket("xxx-img"),
			WithImgUrl("http://xx.xx"),
			WithFilePath("./img/"),
			WithZone(&storage.ZoneHuabei),
			WithIsDelLocal(true),
		)

		qiniu.SetFilePath("./img1/")
		qiniu.SetFileName(fmt.Sprintf("%s%d", "test", time.Now().UnixNano()))

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

	// r.Run(":8080")
}
