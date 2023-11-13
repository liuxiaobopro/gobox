package district

import (
	"testing"
)

func TestDistrict_Query(t *testing.T) {
	d := NewDistrict("xxx",
		WithKeywords("铁西"),
		WithSubdistrict("0"),
		WithExtensions("all"))

	resp, err := d.Query()
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(string(resp))
}
