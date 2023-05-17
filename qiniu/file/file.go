package file

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"
	"sync"

	filex "github.com/liuxiaobopro/gobox/file"
	qiniux "github.com/liuxiaobopro/gobox/qiniu"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

type File struct {
	qiniux.Qiniu
	Bucket string // 空间名
	ImgUrl string // cdn域名

	lock sync.Mutex

	service struct {
		FileName   string // 上传到服务器的文件名
		FilePath   string // 上传到服务器的文件路径
		IsDelLocal bool   // 上传之后是否删除本地文件
		Zone       *storage.Zone
	}
}

type option func(*File)

func WithKey(accessKey, secretKey string) option {
	return func(q *File) {
		q.AccessKey = accessKey
		q.SecretKey = secretKey
	}
}

func WithBucket(bucket string) option {
	return func(q *File) {
		q.Bucket = bucket
	}
}

func WithImgUrl(imgUrl string) option {
	return func(q *File) {
		q.ImgUrl = imgUrl
	}
}

func WithFilePath(filePath string) option {
	return func(q *File) {
		q.service.FilePath = filePath
	}
}

func WithIsDelLocal(isDelLocal bool) option {
	return func(q *File) {
		q.service.IsDelLocal = isDelLocal
	}
}

func WithZone(zone *storage.Zone) option {
	return func(q *File) {
		q.service.Zone = zone
	}
}

func NewQiniu(opts ...option) *File {
	q := &File{}
	for _, opt := range opts {
		opt(q)
	}

	if q.service.FilePath == "" {
		q.service.FilePath = "./"
	}

	if q.service.Zone == nil {
		q.service.Zone = &storage.ZoneHuanan
	}
	return q
}

func (f *File) SetFilePath(filePath string) {
	f.service.FilePath = filePath
}

func (f *File) SetFileName(fileName string) {
	f.service.FileName = fileName
}

func (f *File) UploadFile(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	f.lock.Lock()
	defer f.lock.Unlock()

	filePath, fileName, err := filex.Upload(file, fileHeader, f.service.FilePath, f.service.FileName)
	if err != nil {
		return "", err
	}
	// 简单上传的凭证
	putPolicy := storage.PutPolicy{
		Scope: f.Bucket,
	}
	mac := qbox.NewMac(f.AccessKey, f.SecretKey)
	upToken := putPolicy.UploadToken(mac)

	// 空间对应机房
	// 其中关于Zone对象和机房的关系如下：
	//    华东    storage.ZoneHuadong
	//    华北    storage.ZoneHuabei
	//    华南    storage.ZoneHuanan
	//    北美    storage.ZoneBeimei
	cfg := storage.Config{
		Zone:          f.service.Zone, // 七牛云存储空间设置首页有存储区域
		UseCdnDomains: false,          // 不启用HTTPS域名
		UseHTTPS:      false,          // 不使用CND加速
	}

	// 构建上传表单对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	// 可选
	putExtra := storage.PutExtra{
		// Params: map[string]string{
		// 	"x:name": "github logo",
		// },
	}

	// err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err := formUploader.PutFile(context.Background(), &ret, upToken, fileName, filePath, &putExtra); err != nil {
		return "", err
	}
	url := fmt.Sprintf("%s/%s", f.ImgUrl, ret.Key)

	// 删除本地文件
	if f.service.IsDelLocal {
		defer func(name string) {
			_ = os.Remove(name)
		}(filePath)
	}
	return url, nil
}
