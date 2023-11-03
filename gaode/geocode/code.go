package geocode

import (
	"fmt"

	gaodex "github.com/liuxiaobopro/gobox/gaode"
	httpx "github.com/liuxiaobopro/gobox/http"
)

type GeocodeCodeRes struct {
	Status   string            `json:"status,omitempty"`
	Info     string            `json:"info,omitempty"`
	InfoCode string            `json:"infocode,omitempty"`
	Count    string            `json:"count,omitempty"`
	Geocodes []GeocodeCodeItem `json:"geocodes,omitempty"`
}

type GeocodeCodeItem struct {
	FormattedAddress string      `json:"formatted_address,omitempty"`
	Country          string      `json:"country,omitempty"`
	Province         string      `json:"province,omitempty"`
	CityCode         string      `json:"citycode,omitempty"`
	City             interface{} `json:"city,omitempty"`
	District         interface{} `json:"district,omitempty"`
	Township         []string    `json:"township,omitempty"`
	Neighborhood     struct {
		Name []string `json:"name,omitempty"`
		Type []string `json:"type,omitempty"`
	} `json:"neighborhood,omitempty"`
	Building struct {
		Name []string `json:"name,omitempty"`
		Type []string `json:"type,omitempty"`
	} `json:"building,omitempty"`
	Adcode   string   `json:"adcode,omitempty"`
	Street   []string `json:"street,omitempty"`
	Number   []string `json:"number,omitempty"`
	Location string   `json:"location,omitempty"`
	Level    string   `json:"level,omitempty"`
}

type GeocodeCode struct {
	gaodex.Gaode

	Address string
}

func WithGeocodeCodeAddress(address string) func(*GeocodeCode) {
	return func(g *GeocodeCode) {
		g.Address = address
	}
}

func NewGeocodeCode(key string, options ...func(*GeocodeCode)) *GeocodeCode {
	g := &GeocodeCode{
		Gaode: gaodex.Gaode{
			Key: key,
		},
		Address: "",
	}

	for _, option := range options {
		option(g)
	}

	return g
}

func (g *GeocodeCode) Query() ([]byte, error) {
	client := &httpx.Client{
		Url: g.url(),
	}

	return client.Get()
}

func (g *GeocodeCode) url() string {
	return fmt.Sprintf("https://restapi.amap.com/v3/geocode/geo?address=%s&key=%s", g.Address, g.Key)
}
