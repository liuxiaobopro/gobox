package geocode

import (
	"encoding/json"
	"testing"
)

func TestNewGeocodeRegeo(t *testing.T) {
	key := "xxx"

	geo := NewGeocodeRegeo(key, WithGeocodeRegeoLocation("118.816773,35.415796"))

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
}
