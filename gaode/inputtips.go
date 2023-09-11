package gaode

import (
	"fmt"

	httpx "github.com/liuxiaobopro/gobox/http"
)

type Inputtips struct {
	Gaode

	Keywords string
}

func WithKeywords(keywords string) func(*Inputtips) {
	return func(g *Inputtips) {
		g.Keywords = keywords
	}
}

func NewInputtips(key string, options ...func(*Inputtips)) *Inputtips {
	g := &Inputtips{
		Gaode: Gaode{
			Key: key,
		},
		Keywords: "",
	}

	for _, option := range options {
		option(g)
	}

	return g
}

func (g *Inputtips) Query() ([]byte, error) {
	client := &httpx.Client{
		Url: g.url(),
	}

	return client.Get()
}

func (g *Inputtips) url() string {
	return fmt.Sprintf("https://restapi.amap.com/v3/assistant/inputtips?keywords=%s&key=%s", g.Keywords, g.Key)
}
