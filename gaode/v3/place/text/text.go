package text

import (
	"fmt"

	gaodex "github.com/liuxiaobopro/gobox/gaode"
	httpx "github.com/liuxiaobopro/gobox/http"
)

/*
搜索POI - 关键字搜索

https://lbs.amap.com/api/webservice/guide/api/search#text
*/

type District struct {
	gaodex.Gaode

	Keywords   string
	City       string
	Citylimit  string
	Children   string
	Offset     string
	Page       string
	Extensions string
}

func WithKeywords(keywords string) func(*District) {
	return func(d *District) {
		d.Keywords = keywords
	}
}

func WithCity(city string) func(*District) {
	return func(d *District) {
		d.City = city
	}
}

func WithCitylimit(citylimit string) func(*District) {
	return func(d *District) {
		d.Citylimit = citylimit
	}
}

func WithChildren(children string) func(*District) {
	return func(d *District) {
		d.Children = children
	}
}

func WithOffset(offset string) func(*District) {
	return func(d *District) {
		d.Offset = offset
	}
}

func WithPage(page string) func(*District) {
	return func(d *District) {
		d.Page = page
	}
}

func WithExtensions(extensions string) func(*District) {
	return func(d *District) {
		d.Extensions = extensions
	}
}

func NewText(key string, options ...func(*District)) *District {
	d := &District{
		Gaode: gaodex.Gaode{
			Key: key,
		},
		Keywords:   "",
		City:       "",
		Citylimit:  "",
		Children:   "",
		Offset:     "",
		Page:       "",
		Extensions: "",
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
	baseUrl := "https://restapi.amap.com/v3/place/text"
	return fmt.Sprintf("%s?key=%s&keywords=%s&city=%s&citylimit=%s&children=%s&offset=%s&page=%s&extensions=%s", baseUrl, d.Key, d.Keywords, d.City, d.Citylimit, d.Children, d.Offset, d.Page, d.Extensions)
}
