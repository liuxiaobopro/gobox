package qrcode

import (
	"fmt"
	"testing"
	"time"
)

func TestQrCode_Create(t *testing.T) {
	qr := New(
		"https://baidu.com/",
		fmt.Sprintf("../temp/baidu_%s.png", time.Now().Format("20060102150405")),
		WithMargin(-50),
	)

	if err := qr.Create(); err != nil {
		t.Fatalf("Create qr code failed: %v", err)
	}
}
