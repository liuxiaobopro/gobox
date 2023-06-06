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
	Bucket     string // 空间名
	ImgUrl     string // cdn域名
	ServerPath string // 服务器文件路径

	lock sync.Mutex

	service struct {
		FileName      string // 上传到服务器的文件名
		FilePath      string // 本地文件的文件路径
		IsDelLocal    bool   // 上传之后是否删除本地文件
		Zone          *storage.Zone
		UseHTTPS      bool
		UseCdnDomains bool
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

func WithServerPath(serverPath string) option {
	return func(q *File) {
		q.ServerPath = serverPath
	}
}

func WithUseHTTPS(useHTTPS bool) option {
	return func(q *File) {
		q.service.UseHTTPS = useHTTPS
	}
}

func WithUseCdnDomains(useCdnDomains bool) option {
	return func(q *File) {
		q.service.UseCdnDomains = useCdnDomains
	}
}

func NewFile(opts ...option) *File {
	q := &File{}

	q.service.UseHTTPS = true
	q.service.UseCdnDomains = true

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

	putPolicy := storage.PutPolicy{
		Scope: f.Bucket,
	}
	mac := qbox.NewMac(f.AccessKey, f.SecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Region:        f.service.Zone,
		UseHTTPS:      f.service.UseHTTPS,
		UseCdnDomains: f.service.UseCdnDomains,
	}

	resumeUploader := storage.NewResumeUploaderV2(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.RputV2Extra{}

	if err := resumeUploader.PutFile(context.Background(), &ret, upToken, f.ServerPath+fileName, filePath, &putExtra); err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/%s", f.ImgUrl, ret.Key)

	if f.service.IsDelLocal {
		defer func(name string) {
			_ = os.Remove(name)
		}(filePath)
	}

	return url, nil
}
