package lbs

import "math"

type Ucs string // 坐标系

const (
	WGS84 Ucs = "wgs84" // WGS84坐标系
	GCJ02 Ucs = "gcj02" // 国测局坐标系
	BD09  Ucs = "bd09"  // 百度坐标系
)

type LngLat struct {
	Lng float64 // 经度
	Lat float64 // 纬度
}

type Point struct {
	LngLat
	From Ucs // 坐标系
}

type Result struct {
	WGS84 *LngLat
	GCJ02 *LngLat
	BD09  *LngLat
}

const (
	xPi = 3.14159265358979324 * 3000.0 / 180.0
	a   = 6378245.0
	ee  = 0.00669342162296594323
)

func outOfChina(lng, lat float64) bool {
	if lng < 72.004 || lng > 137.8347 {
		return true
	}
	if lat < 0.8293 || lat > 55.8271 {
		return true
	}
	return false
}

func transformLat(lng, lat float64) float64 {
	ret := -100.0 + 2.0*lng + 3.0*lat + 0.2*lat*lat + 0.1*lng*lat + 0.2*math.Sqrt(math.Abs(lng))
	ret += (20.0*math.Sin(6.0*lng*math.Pi) + 20.0*math.Sin(2.0*lng*math.Pi)) * 2.0 / 3.0
	ret += (20.0*math.Sin(lat*math.Pi) + 40.0*math.Sin(lat/3.0*math.Pi)) * 2.0 / 3.0
	ret += (160.0*math.Sin(lat/12.0*math.Pi) + 320*math.Sin(lat*math.Pi/30.0)) * 2.0 / 3.0
	return ret
}

func transformLng(lng, lat float64) float64 {
	ret := 300.0 + lng + 2.0*lat + 0.1*lng*lng + 0.1*lng*lat + 0.1*math.Sqrt(math.Abs(lng))
	ret += (20.0*math.Sin(6.0*lng*math.Pi) + 20.0*math.Sin(2.0*lng*math.Pi)) * 2.0 / 3.0
	ret += (20.0*math.Sin(lng*math.Pi) + 40.0*math.Sin(lng/3.0*math.Pi)) * 2.0 / 3.0
	ret += (150.0*math.Sin(lng/12.0*math.Pi) + 300.0*math.Sin(lng/30.0*math.Pi)) * 2.0 / 3.0
	return ret
}

func NewPoint(lng, lat float64, from Ucs) *Point {
	return &Point{
		LngLat: LngLat{
			Lng: lng,
			Lat: lat,
		},
		From: from,
	}
}

func (p *Point) Convert() *Result {
	result := &Result{}
	switch p.From {
	case WGS84:
		result.WGS84 = &p.LngLat
		result.GCJ02 = wgs84ToGcj02(&p.LngLat)
		result.BD09 = wgs84ToBd09(&p.LngLat)
	case GCJ02:
		result.GCJ02 = &p.LngLat
		result.WGS84 = gcj02ToWgs84(&p.LngLat)
		result.BD09 = gcj02ToBd09(&p.LngLat)
	case BD09:
		result.BD09 = &p.LngLat
		result.WGS84 = bd09ToWgs84(&p.LngLat)
		result.GCJ02 = bd09ToGcj02(&p.LngLat)
	}
	return result
}

func wgs84ToGcj02(p *LngLat) *LngLat {
	lng := p.Lng
	lat := p.Lat

	if outOfChina(lng, lat) {
		return &LngLat{Lng: lng, Lat: lat}
	}

	dlat := transformLat(lng-105.0, lat-35.0)
	dlng := transformLng(lng-105.0, lat-35.0)
	radlat := lat / 180.0 * math.Pi
	magic := math.Sin(radlat)
	magic = 1 - ee*magic*magic
	sqrtmagic := math.Sqrt(magic)
	dlat = (dlat * 180.0) / ((a * (1 - ee)) / (magic * sqrtmagic) * math.Pi)
	dlng = (dlng * 180.0) / (a / sqrtmagic * math.Cos(radlat) * math.Pi)

	mglat := lat + dlat
	mglng := lng + dlng

	return &LngLat{Lng: mglng, Lat: mglat}
}

func wgs84ToBd09(p *LngLat) *LngLat {
	gcj02 := wgs84ToGcj02(p)
	return gcj02ToBd09(gcj02)
}

func gcj02ToWgs84(p *LngLat) *LngLat {
	lng := p.Lng
	lat := p.Lat

	if outOfChina(lng, lat) {
		return &LngLat{Lng: lng, Lat: lat}
	}

	dlat := transformLat(lng-105.0, lat-35.0)
	dlng := transformLng(lng-105.0, lat-35.0)
	radlat := lat / 180.0 * math.Pi
	magic := math.Sin(radlat)
	magic = 1 - ee*magic*magic
	sqrtmagic := math.Sqrt(magic)
	dlat = (dlat * 180.0) / ((a * (1 - ee)) / (magic * sqrtmagic) * math.Pi)
	dlng = (dlng * 180.0) / (a / sqrtmagic * math.Cos(radlat) * math.Pi)

	mglat := lat + dlat
	mglng := lng + dlng

	return &LngLat{Lng: lng*2 - mglng, Lat: lat*2 - mglat}
}

func gcj02ToBd09(p *LngLat) *LngLat {
	lng := p.Lng
	lat := p.Lat

	z := math.Sqrt(lng*lng+lat*lat) + 0.00002*math.Sin(lat*xPi)
	theta := math.Atan2(lat, lng) + 0.000003*math.Cos(lng*xPi)
	bdLng := z*math.Cos(theta) + 0.0065
	bdLat := z*math.Sin(theta) + 0.006

	return &LngLat{Lng: bdLng, Lat: bdLat}
}

func bd09ToWgs84(p *LngLat) *LngLat {
	gcj02 := bd09ToGcj02(p)
	return gcj02ToWgs84(gcj02)
}

func bd09ToGcj02(p *LngLat) *LngLat {
	lng := p.Lng - 0.0065
	lat := p.Lat - 0.006
	z := math.Sqrt(lng*lng+lat*lat) - 0.00002*math.Sin(lat*xPi)
	theta := math.Atan2(lat, lng) - 0.000003*math.Cos(lng*xPi)
	gcjLng := z * math.Cos(theta)
	gcjLat := z * math.Sin(theta)

	return &LngLat{Lng: gcjLng, Lat: gcjLat}
}
