package sms

import (
	"testing"
)

func TestNewSms(t *testing.T) {
	sms := NewSms(
		WithAssessKey("xxx", "xxx"),
		WithSignName("xxx"),
		WithTemplateCode("xxxx"),
		WithTemplateParam("{\"code\":\"%s\"}"),
		WithConnectTimeout(5),
	)

	if err := sms.CreateClient(); err != nil {
		t.Error(err)
	}
	sms.SetPhone("1xxxx2")
	sms.SetCode("2345")
	if res, err := sms.Send(); err != nil {
		t.Error(err)
	} else {
		t.Logf("send result: %#v", res)
	}
}
