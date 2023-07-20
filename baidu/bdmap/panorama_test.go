package bdmap

import (
	"testing"
)

func TestDbMap_Panorama(t *testing.T) {
	dbPanorama := NewPanorama("xxx",
		WithWidth(512),
		WithHeight(256),
		WithLocation("116.313393,40.04778"),
		WithFov(180),
	)

	dbPanorama.Panorama()
}
