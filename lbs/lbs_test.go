package lbs

import (
	"fmt"
	"testing"
)

func TestNewPoint(t *testing.T) {
	point := &Point{
		LngLat: LngLat{
			Lng: 122.86042300000000,
			Lat: 41.69003600000000,
		},
		From: WGS84,
	}

	res := point.Convert()

	fmt.Println("WGS84: ", res.WGS84)
	fmt.Println("GCJ02: ", res.GCJ02)
	fmt.Println("BD09: ", res.BD09)
}
