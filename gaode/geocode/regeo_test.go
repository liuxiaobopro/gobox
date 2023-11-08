package geocode

import (
	"encoding/json"
	"testing"
)

func TestNewGeocodeRegeo(t *testing.T) {
	key := "xxx"

	geo := NewGeocodeRegeo(key, WithGeocodeRegeoLocation("118.2012,39.6463666666667"))

	b, err := geo.Query()

	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%s", b)

	var res GeocodeRegeoRes
	err = json.Unmarshal(b, &res)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%s", res.Regeocode.AddressComponent.Province)
	t.Logf("%s", res.Regeocode.AddressComponent.City)
	t.Logf("%s", res.Regeocode.AddressComponent.District)
}
