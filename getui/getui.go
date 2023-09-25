package getui

const (
	BaseUrlV2 = "https://restapi.getui.com/v2"

	// 鉴权api
	AuthUrl = BaseUrlV2 + "/%s/auth"
)

type Config struct {
	AppId        string
	AppKey       string
	AppSecret    string
	MasterSecret string

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
