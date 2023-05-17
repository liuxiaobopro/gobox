package sms

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/liuxiaobopro/gobox/alibaba"
)

type Sms struct {
	alibaba.Alibaba
	Sms struct {
		SignName       string // 短信签名名称
		TemplateCode   string // 短信模板ID
		TemplateParam  string // 短信模板变量对应的实际值，JSON格式
		ConnectTimeout int    // 连接超时时间
	}
	service struct {
		client     *dysmsapi20170525.Client
		phone      string // 手机号
		sendResult *SendResult
	}
}

type option func(*Sms)

func WithAssessKey(accessKeyId string, accessKeySecret string) option {
	return func(s *Sms) {
		s.AccessKeyId = accessKeyId
		s.AccessKeySecret = accessKeySecret
	}
}

func WithSignName(signName string) option {
	return func(s *Sms) {
		s.Sms.SignName = signName
	}
}

func WithTemplateCode(templateCode string) option {
	return func(s *Sms) {
		s.Sms.TemplateCode = templateCode
	}
}

func WithTemplateParam(templateParam string) option {
	return func(s *Sms) {
		s.Sms.TemplateParam = templateParam
	}
}

func WithConnectTimeout(connectTimeout int) option {
	return func(s *Sms) {
		s.Sms.ConnectTimeout = connectTimeout
	}
}

func WithPhone(phone string) option {
	return func(s *Sms) {
		s.service.phone = phone
	}
}

func NewSms(options ...option) *Sms {
	s := &Sms{}
	for _, option := range options {
		option(s)
	}
	return s
}

type SendResult struct {
	*dysmsapi20170525.SendSmsResponse
}

func (s *Sms) CreateClient() error {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: tea.String(s.AccessKeyId),
		// 必填，您的 AccessKey Secret
		AccessKeySecret: tea.String(s.AccessKeySecret),
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result := &dysmsapi20170525.Client{}
	_result, _err := dysmsapi20170525.NewClient(config)
	if _err != nil {
		return _err
	}
	s.service.client = _result
	return nil
}

// Send 发送短信
func (s *Sms) Send() error {
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  tea.String(s.service.phone),
		SignName:      tea.String(s.Sms.SignName),
		TemplateCode:  tea.String(s.Sms.TemplateCode),
		TemplateParam: tea.String(s.Sms.TemplateParam),
	}
	runtime := &util.RuntimeOptions{
		ConnectTimeout: tea.Int(s.Sms.ConnectTimeout),
	}
	_result, _err := s.service.client.SendSmsWithOptions(sendSmsRequest, runtime)
	if _err != nil {
		return _err
	}
	s.service.sendResult = &SendResult{_result}
	return nil
}

func (s *Sms) SendResult() *SendResult {
	return s.service.sendResult
}

func (s *Sms) Ok() bool {
	return *s.service.sendResult.Body.Code == "OK"
}
