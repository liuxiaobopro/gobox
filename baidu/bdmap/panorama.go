package bdmap

import (
	"encoding/json"

	httpx "github.com/liuxiaobopro/gobox/http"
)

type Panorama struct {
	DbMap

	Mcode     string `json:"mcode"`     // 安全码。若为Android/IOS SDK的ak, 该参数必需。
	Width     int    `json:"width"`     // 图片宽度，范围[10,1024]
	Height    int    `json:"height"`    // 图片高度，范围[10,512]
	Location  string `json:"location"`  // 全景位置点坐标。坐标格式：lng<经度>，lat<纬度>，例如116.313393,40.047783。
	Coordtype string `json:"coordtype"` // 全景位置点坐标。坐标格式：lng<经度>，lat<纬度>，例如116.313393,40.047783。
	Poiid     string `json:"poiid"`     // poi的id，该属性通常通过place api接口获取，poiid与panoid、location一起设置全景的显示场景，优先级为：poiid>panoid>location。其中根据poiid获取的全景视角最佳。
	Panoid    string `json:"panoid"`    // 全景图id，panoid与poiid、location一起设置全景的显示场景，优先级为：poiid>panoid>location。
	Heading   int    `json:"heading"`   // 水平视角，范围[0,360]
	Pitch     int    `json:"pitch"`     // 垂直视角，范围[0,90]。
	Fov       int    `json:"fov"`       // 水平方向范围，范围[10,360]，fov=360即可显示整幅全景图
}

type PanoramaOption func(*Panorama)

func WithMcode(mcode string) PanoramaOption {
	return func(m *Panorama) {
		m.Mcode = mcode
	}
}

func WithWidth(width int) PanoramaOption {
	return func(m *Panorama) {
		m.Width = width
	}
}

func WithHeight(height int) PanoramaOption {
	return func(m *Panorama) {
		m.Height = height
	}
}

func WithLocation(location string) PanoramaOption {
	return func(m *Panorama) {
		m.Location = location
	}
}

func WithCoordtype(coordtype string) PanoramaOption {
	return func(m *Panorama) {
		m.Coordtype = coordtype
	}
}

func WithPoiid(poiid string) PanoramaOption {
	return func(m *Panorama) {
		m.Poiid = poiid
	}
}

func WithPanoid(panoid string) PanoramaOption {
	return func(m *Panorama) {
		m.Panoid = panoid
	}
}

func WithHeading(heading int) PanoramaOption {
	return func(m *Panorama) {
		m.Heading = heading
	}
}

func WithPitch(pitch int) PanoramaOption {
	return func(m *Panorama) {
		m.Pitch = pitch
	}
}

func WithFov(fov int) PanoramaOption {
	return func(m *Panorama) {
		m.Fov = fov
	}
}

type PanoramaReply struct{}

var (
	panoramaRouter = "/panorama/v2" // 全景图api路由

	panoramaUrl = baseUrl + panoramaRouter
)

func NewPanorama(ak string, opts ...PanoramaOption) *Panorama {
	m := &Panorama{
		DbMap: DbMap{
			Ak: ak,
		},
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// Panorama 全景图 (文档: https://lbs.baidu.com/faq/api?title=viewstatic-base)
func (th *Panorama) Panorama() (interface{}, error) {
	bb, _ := json.Marshal(th)

	var m map[string]string
	_ = json.Unmarshal(bb, &m)

	httpCli := &httpx.Client{
		Url:    panoramaUrl,
		Params: m,
	}

	b, err := httpCli.Get()
	if err != nil {
		return nil, err
	}

	return b, nil
}
