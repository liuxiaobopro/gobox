package gaode

import (
	"fmt"

	httpx "github.com/liuxiaobopro/gobox/http"
)

type Geocode struct {
	Gaode

	Address string
}

func WithAddress(address string) func(*Geocode) {
	return func(g *Geocode) {
		g.Address = address
	}
}

func NewGeocode(key string, options ...func(*Geocode)) *Geocode {
	g := &Geocode{
		Gaode: Gaode{
			Key: key,
		},
		Address: "",
	}

	for _, option := range options {
		option(g)
	}

	return g
}

func (g *Geocode) Query() ([]byte, error) {
	client := &httpx.Client{
		Url: g.url(),
	}

	return client.Get()
}

func (g *Geocode) url() string {
	return fmt.Sprintf("https://restapi.amap.com/v3/geocode/geo?address=%s&key=%s", g.Address, g.Key)
}
