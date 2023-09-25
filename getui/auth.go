package getui

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"

	httpx "github.com/liuxiaobopro/gobox/http"
	logx "github.com/liuxiaobopro/gobox/log"
	timex "github.com/liuxiaobopro/gobox/time"
)

type AuthT struct {
	Sign      string `json:"sign"`
	Timestamp string `json:"timestamp"`
	Appkey    string `json:"appkey"`
}

func (th *Config) Auth() error {
	/*
		鉴权
		文档地址: https://docs.getui.com/getui/server/rest_v2/token/
	*/

	timestamp := strconv.FormatInt(timex.NowMilliTimeStamp(), 10)
	str := th.AppKey + timestamp + th.MasterSecret

	// sha256 加密
	hash := sha256.New()
	hash.Write([]byte(str))
	sum := hash.Sum(nil)

	auth := &AuthT{
		Sign:      fmt.Sprintf("%x", sum),
		Timestamp: timestamp,
		Appkey:    th.AppKey,
	}

	jsonBytes, err := json.Marshal(auth)
	if err != nil {
		logx.Errorf("鉴权失败: %v", err)
		return err
	}

	hc := &httpx.Client{
		Url: fmt.Sprintf(AuthUrl, th.AppId),
		Header: map[string]string{
			"Content-Type": "application/json",
			"Charset":      "UTF-8",
		},
		Json: jsonBytes,
	}

	res, err := hc.Post()
	if err != nil {
		return err
	}

	var r *AuthReply
	if err := json.Unmarshal(res, &r); err != nil {
		return err
	}

	if r.Code != 0 {
		return fmt.Errorf("鉴权失败: %s", r.Msg)
	}

	th.ExpireTime = r.Data.ExpireTime
	th.Token = r.Data.Token

	return nil
}
