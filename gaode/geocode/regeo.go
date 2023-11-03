package geocode

import (
	"fmt"

	gaodex "github.com/liuxiaobopro/gobox/gaode"
	httpx "github.com/liuxiaobopro/gobox/http"
)

type AddressComponent struct {
	Province string `json:"province,omitempty"`
}

type Regeocode struct {
	AddressComponent AddressComponent `json:"addressComponent,omitempty"`
	FormattedAddress string           `json:"formatted_address,omitempty"`
}

type GeocodeRegeoRes struct {
	Status    string    `json:"status,omitempty"`
	Regeocode Regeocode `json:"regeocode,omitempty"`
	Info      string    `json:"info,omitempty"`
	Infocode  string    `json:"infocode,omitempty"`
}

type GeocodeRegeo struct {
	gaodex.Gaode

	Location string
}

func WithGeocodeRegeoLocation(location string) func(*GeocodeRegeo) {
	return func(g *GeocodeRegeo) {
		g.Location = location
	}
}

func NewGeocodeRegeo(key string, options ...func(*GeocodeRegeo)) *GeocodeRegeo {
	g := &GeocodeRegeo{
		Gaode: gaodex.Gaode{
			Key: key,
		},
		Location: "",
	}

	for _, option := range options {
		option(g)
	}

	return g
}

func (g *GeocodeRegeo) Query() ([]byte, error) {
	client := &httpx.Client{
		Url: g.url(),
	}

	return client.Get()
}

func (g *GeocodeRegeo) url() string {
	return fmt.Sprintf("https://restapi.amap.com/v3/geocode/regeo?location=%s&key=%s", g.Location, g.Key)
}
