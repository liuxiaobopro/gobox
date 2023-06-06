package file

import (
	"bytes"
	"context"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

func TestAppend(t *testing.T) {
	type args struct {
		path    string
		content string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test1", args{"./test.txt", "test\n"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Append(tt.args.path, tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("Append() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHas(t *testing.T) {
	type args struct {
		path    string
		content string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"test1", args{"./test.txt", "test"}, true, false},
		{"test2", args{"./test.txt", "test1"}, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Has(tt.args.path, tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("Has() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Has() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindStringsBetween(t *testing.T) {
	str := `
    package file
    import (
        "bufio"
        "io/ioutil"
        "os"
        "path/filepath"
        "regexp"
        "strings"
    )
    // ReplaceInDir 替换目录下所有文件中的字符串
    `
	reg := "import \\((?s)(.*?)\\)"
	got := FindStringsBetween(str, reg)
	t.Log("got: ", got)
}

func TestUploadChunkFile(t *testing.T) {
	var ak = "7NPqgfSKvc2b8-vg9EA7o18eTHa3qGsCFaOkq8DY"
	var sk = "YekfJPpdi5J84TZfAZbSHOQWVsTcjXxIfBkMuGKQ"
	var bucket = "yinpinshenyang"
	var url = "https://qiniu.xuanwh.com/"

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(ak, sk)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuabei,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	resumeUploaderV2 := storage.NewResumeUploaderV2(&cfg)
	upHost, err := resumeUploaderV2.UpHost(ak, bucket)
	if err != nil {
		t.Fatal(err)
	}
	key := "go-cloud-storage/lala.mp4"
	// 初始化分块上传
	initPartsRet := storage.InitPartsRet{}
	err = resumeUploaderV2.InitParts(context.TODO(), upToken, upHost, bucket, key, true, &initPartsRet)
	if err != nil {
		t.Fatal(err)
	}

	fileInfo, err := os.Open("./ASBJVID.mp4")
	if err != nil {
		t.Fatal(err)
	}
	defer fileInfo.Close()
	fileContent, err := ioutil.ReadAll(fileInfo)
	if err != nil {
		t.Fatal(err)
	}
	fileLen := len(fileContent)
	chunkSize2 := 2 * 1024 * 1024

	num := fileLen / chunkSize2
	if fileLen%chunkSize2 > 0 {
		num++
	}

	// 分块上传
	var uploadPartInfos []storage.UploadPartInfo
	for i := 1; i <= num; i++ {
		partNumber := int64(i)
		fmt.Printf("开始上传第%v片数据", partNumber)

		var partContentBytes []byte
		endSize := i * chunkSize2
		if endSize > fileLen {
			endSize = fileLen
		}
		partContentBytes = fileContent[(i-1)*chunkSize2 : endSize]
		partContentMd5 := fmt.Sprintf("%x", md5.Sum(partContentBytes))
		uploadPartsRet := storage.UploadPartsRet{}
		err = resumeUploaderV2.UploadParts(context.TODO(), upToken, upHost, bucket, key, true,
			initPartsRet.UploadID, partNumber, partContentMd5, &uploadPartsRet, bytes.NewReader(partContentBytes),
			len(partContentBytes))
		if err != nil {
			t.Fatal(err)
		}
		uploadPartInfos = append(uploadPartInfos, storage.UploadPartInfo{
			Etag:       uploadPartsRet.Etag,
			PartNumber: partNumber,
		})
		fmt.Printf("结束上传第%d片数据\n", partNumber)
	}

	// 完成上传
	rPutExtra := storage.RputV2Extra{Progresses: uploadPartInfos}
	comletePartRet := storage.PutRet{}
	err = resumeUploaderV2.CompleteParts(context.TODO(), upToken, upHost, &comletePartRet, bucket, key,
		true, initPartsRet.UploadID, &rPutExtra)
	if err != nil {
		t.Fatal(err)
	}

	url2 := url + comletePartRet.Key
	// fmt.Println(comletePartRet.Hash)
	fmt.Println(url2)
}
