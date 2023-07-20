package bdmap

type DbMap struct {
	Ak string `json:"ak"` // 用户的访问密钥。支持浏览器端和服务端ak，网页应用推荐使用服务端ak
}

type DbMapOption func(*DbMap)

var (
	baseUrl = "https://api.map.baidu.com" // 百度地图api基础url
)

func NewDbMap(ak string, opts ...DbMapOption) *DbMap {
	m := &DbMap{
		Ak: ak,
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}
