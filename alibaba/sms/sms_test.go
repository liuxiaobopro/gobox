package sms

import (
	"testing"
)

func TestNewSms(t *testing.T) {
	sms := NewSms(
		WithAssessKey("xxx", "xxxx"),
		WithSignName("xxx"),
		WithTemplateCode("xxxx"),
		WithTemplateParam(`{"code":"123456"}`),
		WithConnectTimeout(5),
		WithPhone("xxxx"),
	)

	if err := sms.CreateClient(); err != nil {
		t.Error(err)
	}
	if err := sms.Send(); err != nil {
		t.Error(err)
	}

	t.Logf("send result: %#v", sms.SendResult())
	if sms.Ok() {
		t.Log("send ok")
	} else {
		t.Log("send failed")
	}
}
