package geocode

import (
	"testing"
)

func TestGeocode_Query(t *testing.T) {
	key := "xxx"

	geo := NewGeocodeCode(key, WithGeocodeCodeAddress("北京市朝阳区阜通东大街6号"))

	b, err := geo.Query()

	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%s", b)
}
