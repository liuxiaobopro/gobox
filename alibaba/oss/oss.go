package oss

import (
	"io"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/liuxiaobopro/gobox/alibaba"
)

type Oss struct {
	alibaba.Alibaba
	Oss struct {
		Endpoint string // 地域节点
		Bucket   string // 存储空间名称
	}
	service struct {
		client *oss.Client
	}
}

func WithAssessKey(accessKeyId string, accessKeySecret string) option {
	return func(s *Oss) {
		s.AccessKeyId = accessKeyId
		s.AccessKeySecret = accessKeySecret
	}
}

func WithEndpoint(endpoint string) option {
	return func(s *Oss) {
		s.Oss.Endpoint = endpoint
	}
}

func WithBucket(bucket string) option {
	return func(s *Oss) {
		s.Oss.Bucket = bucket
	}
}

type option func(*Oss)

func NewOss(options ...option) *Oss {
	s := &Oss{}
	for _, option := range options {
		option(s)
	}
	return s
}

// CreateClient 创建客户端
func (s *Oss) CreateClient() error {
	var (
		client *oss.Client
		err    error
	)
	client, err = oss.New(s.Oss.Endpoint, s.AccessKeyId, s.AccessKeySecret)
	if err != nil {
		return err
	}
	s.service.client = client
	return nil
}

// PutObject 上传文件
func (s *Oss) PutObject(objectKey string, localFile io.Reader) error {
	bucket, err := s.service.client.Bucket(s.Oss.Bucket)
	if err != nil {
		return err
	}
	err = bucket.PutObject(objectKey, localFile)
	if err != nil {
		return err
	}
	return nil
}
