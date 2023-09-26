package getui

/*
部分代码参考: https://github.com/dacker-soul/getui
*/

const (
	BaseUrlV2 = "https://restapi.getui.com/v2/%s"

	// 鉴权api
	AuthUrl = "/auth"

	// 【toSingle】执行cid单推
	PushSingleByCidUrl = "/push/single/cid"
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
