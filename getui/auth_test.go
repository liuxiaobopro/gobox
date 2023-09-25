package getui

import (
	"testing"
)

func TestConfig_Auth(t *testing.T) {
	getui := New(
		WithAppId("xxx"),
		WithAppKey("xxx"),
		WithMasterSecret("xxx"),
	)

	if err := getui.Auth(); err != nil {
		t.Error(err)
	}

	t.Log(getui.Token)
}
