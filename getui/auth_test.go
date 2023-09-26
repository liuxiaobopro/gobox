package getui

import (
	"strconv"
	"testing"
	"time"

	timex "github.com/liuxiaobopro/gobox/time"
	"github.com/spf13/viper"
)

func TestConfig_Auth(t *testing.T) {
	//#region 获取配置
	v := viper.New()
	v.SetConfigFile("./config.yaml")
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		t.Error(err)
		return
	}

	var conf ConfigDemo
	if err := v.Unmarshal(&conf); err != nil {
		t.Error(err)
		return
	}
	//#endregion

	//#region 鉴权
	getui := New(
		WithAppId(conf.AppId),
		WithAppKey(conf.AppKey),
		WithMasterSecret(conf.MasterSecret),
	)

	if err := getui.Auth(); err != nil {
		t.Error(err)
		return
	}
	//#endregion

	t.Log(getui.Token)
}

func TestConfig_CheckToken(t *testing.T) {
	//#region 获取配置
	v := viper.New()
	v.SetConfigFile("./config.yaml")
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		t.Error(err)
		return
	}

	var conf ConfigDemo
	if err := v.Unmarshal(&conf); err != nil {
		t.Error(err)
		return
	}
	//#endregion

	//#region 鉴权
	getui := New(
		WithAppId(conf.AppId),
		WithAppKey(conf.AppKey),
		WithMasterSecret(conf.MasterSecret),
	)

	if err := getui.Auth(); err != nil {
		t.Error(err)
		return
	}

	if err := getui.CheckToken(); err != nil {
		t.Error(err)
		return
	}
	//#endregion

	t.Log(getui.ExpireTime)

	et, _ := strconv.ParseInt(getui.ExpireTime, 10, 64)
	t.Log(time.Unix(et, 0).Format(timex.FormatDateTime))
}
