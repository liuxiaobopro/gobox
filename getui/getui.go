package getui

import (
	"encoding/json"

	httpx "github.com/liuxiaobopro/gobox/http"
	logx "github.com/liuxiaobopro/gobox/log"
)

/*
部分代码参考: https://github.com/dacker-soul/getui
*/

const (
	BaseUrlV2 = "https://restapi.getui.com/v2/%s"

	// 鉴权api
	AuthUrl = "/auth"

	// 【toSingle】执行cid单推
	PushSingleByCidUrl = "/push/single/cid"

	// 【toApp】执行群推
	PushAppUrl = "/push/all"
)

type ConfigDemo struct {
	AppId        string
	AppKey       string
	MasterSecret string
	IosCid       string
	AndroidCid   string
}

type Config struct {
	AppId        string
	AppKey       string
	AppSecret    string
	MasterSecret string
	BaseUrl      string

	ExpireTime string
	Token      string
}

func (th *Config) doPost(api string, param interface{}) ([]byte, error) {
	jsonByte, err := json.Marshal(param)
	if err != nil {
		logx.Errorf("json.Marshal err: %v", err)
		return nil, err
	}

	client := &httpx.Client{
		Url: th.BaseUrl + api,
		Header: map[string]string{
			"Content-Type": "application/json",
			"Charset":      "UTF-8",
			"token":        th.Token,
		},
		Json: jsonByte,
	}

	return client.Post()
}

type ConfigOption func(c *Config)

func WithAppId(appId string) ConfigOption {
	return func(c *Config) {
		c.AppId = appId
	}
}

func WithAppKey(appKey string) ConfigOption {
	return func(c *Config) {
		c.AppKey = appKey
	}
}

func WithAppSecret(appSecret string) ConfigOption {
	return func(c *Config) {
		c.AppSecret = appSecret
	}
}

func WithMasterSecret(masterSecret string) ConfigOption {
	return func(c *Config) {
		c.MasterSecret = masterSecret
	}
}

func New(opts ...ConfigOption) *Config {
	c := &Config{}
	for _, opt := range opts {
		opt(c)
	}

	return c
}
