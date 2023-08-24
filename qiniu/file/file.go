package file

import (
	"bytes"
	"context"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"sort"
	"sync"
	"time"

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
		Debug         bool          // 是否开启调试模式
		FileName      string        // 上传到服务器的文件名
		FilePath      string        // 本地文件的文件路径
		IsDelLocal    bool          // 上传之后是否删除本地文件
		Zone          *storage.Zone // 机房
		UseHTTPS      bool          // 是否使用https域名
		UseCdnDomains bool          // 是否使用cdn域名
		ChunkSize     int           // 分片上传的块大小
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

func WithChunkSize(chunkSize int) option {
	return func(q *File) {
		q.service.ChunkSize = chunkSize
	}
}

func WithDebug(debug bool) option {
	return func(q *File) {
		q.service.Debug = debug
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

	if q.service.ChunkSize == 0 {
		q.service.ChunkSize = 10 * 1024 * 1024
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
	upHost, err := resumeUploader.UpHost(f.AccessKey, f.Bucket)
	if err != nil {
		return "", err
	}

	// 初始化分块上传
	key := f.ServerPath + fileName
	initPartsRet := storage.InitPartsRet{}
	if err := resumeUploader.InitParts(context.TODO(), upToken, upHost, f.Bucket, key, true, &initPartsRet); err != nil {
		return "", err
	}

	fileInfo, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer fileInfo.Close()

	fileContent, err := ioutil.ReadAll(fileInfo)
	if err != nil {
		return "", err
	}
	fileLen := len(fileContent)
	chunkSize2 := f.service.ChunkSize

	num := fileLen / chunkSize2
	if fileLen%chunkSize2 > 0 {
		num++
	}

	if f.service.Debug {
		fmt.Println("总共分成", num, "片")
	}
	// 分块上传
	var uploadPartInfos = make([]storage.UploadPartInfo, num)
	var wg sync.WaitGroup
	wg.Add(num)

	for i := 1; i <= num; i++ {
		partNumber := int64(i)

		var partContentBytes []byte
		endSize := i * chunkSize2
		if endSize > fileLen {
			endSize = fileLen
		}
		partContentBytes = fileContent[(i-1)*chunkSize2 : endSize]
		partContentMd5 := fmt.Sprintf("%x", md5.Sum(partContentBytes))
		go func(partNumber int64, partContentBytes []byte) {
			defer wg.Done()

			if f.service.Debug {
				fmt.Printf("开始上传第%v片数据\n", partNumber)
			}

			st := time.Now()
			uploadPartsRet := storage.UploadPartsRet{}
			if err := resumeUploader.UploadParts(context.TODO(), upToken, upHost, f.Bucket, key, true,
				initPartsRet.UploadID, partNumber, partContentMd5, &uploadPartsRet, bytes.NewReader(partContentBytes),
				len(partContentBytes)); err != nil {
				fmt.Printf("上传第%d片数据出错：%v\n", partNumber, err)
				return
			}

			uploadPartInfos[partNumber-1] = storage.UploadPartInfo{
				Etag:       uploadPartsRet.Etag,
				PartNumber: partNumber,
			}

			et := time.Now()
			l := et.Sub(st)

			if f.service.Debug {
				fmt.Printf("上传第%d片数据成功，耗时：%v\n", partNumber, l)
				// fmt.Printf("完成上传第%d片数据\n", partNumber)
			}
		}(partNumber, partContentBytes)
	}
	wg.Wait()

	if f.service.Debug {
		fmt.Println("等待所有分片上传完成")
	}

	// 对分块上传结果进行排序
	sort.Slice(uploadPartInfos, func(i, j int) bool {
		return uploadPartInfos[i].PartNumber < uploadPartInfos[j].PartNumber
	})

	// 完成上传
	rPutExtra := storage.RputV2Extra{Progresses: uploadPartInfos}
	comletePartRet := storage.PutRet{}
	err = resumeUploader.CompleteParts(context.TODO(), upToken, upHost, &comletePartRet, f.Bucket, key,
		true, initPartsRet.UploadID, &rPutExtra)
	if err != nil {
		return "", err
	}

	url := f.ImgUrl + "/" + comletePartRet.Key
	return url, nil
}
