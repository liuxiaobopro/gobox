package mapbox

import (
	"testing"
)

func TestGeocoding_Query(t *testing.T) {
	at := "pk.xxx.xxx"
	gc := NewGeocoding(at,
		WithCountry(CountryGlobal),
		WithLanguage(LanguageChinese),
		WithQuery("白宫"))

	b, err := gc.Query()
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%s", b)
}
