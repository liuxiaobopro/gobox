package district

import (
	"fmt"

	gaodex "github.com/liuxiaobopro/gobox/gaode"
	httpx "github.com/liuxiaobopro/gobox/http"
)

/*
行政区域查询

文档地址: https://lbs.amap.com/api/webservice/guide/api/district
*/

type District struct {
	gaodex.Gaode

	Keywords    string
	Subdistrict string
	Extensions  string
}

func WithKeywords(keywords string) func(*District) {
	return func(d *District) {
		d.Keywords = keywords
	}
}

func WithSubdistrict(subdistrict string) func(*District) {
	return func(d *District) {
		d.Subdistrict = subdistrict
	}
}

func WithExtensions(extensions string) func(*District) {
	return func(d *District) {
		d.Extensions = extensions
	}
}

func NewDistrict(key string, options ...func(*District)) *District {
	d := &District{
		Gaode: gaodex.Gaode{
			Key: key,
		},
		Keywords:    "",
		Subdistrict: "",
		Extensions:  "",
	}

	for _, option := range options {
		option(d)
	}

	return d
}

func (d *District) Query() ([]byte, error) {
	client := &httpx.Client{
		Url: d.url(),
	}

	return client.Get()
}

func (d *District) url() string {
	return fmt.Sprintf("https://restapi.amap.com/v3/config/district?key=%s&keywords=%s&subdistrict=%s&extensions=%s", d.Key, d.Keywords, d.Subdistrict, d.Extensions)
}
