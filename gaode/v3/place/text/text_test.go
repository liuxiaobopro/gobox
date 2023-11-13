package text

import (
	"testing"
)

func TestNewDistrict(t *testing.T) {
	r := NewDistrict("xxx",
		WithKeywords("铁西"),
		WithCity("沈阳"),
		WithCitylimit("true"),
		WithChildren("1"),
		WithOffset("20"),
		WithPage("1"),
		WithExtensions("base"),
	)

	resp, err := r.Query()
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(string(resp))
}
