package getui

import (
	"encoding/json"
	"strconv"
	"testing"
	"time"

	timex "github.com/liuxiaobopro/gobox/time"
	"github.com/spf13/viper"
)

func TestConfig_ToSingleByCid1(t *testing.T) {
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

	//#region 具体业务
	iosChannel := IosChannel{
		Type: "",
		Aps: &Aps{
			Alert: &Alert{
				Title: "卡是谁1-" + time.Now().Format(timex.FormatDateTime),
				Body:  "为什么我们每天都要打 TA ？",
			},
			ContentAvailable: 0,
		},
		AutoBadge:      "+1",
		PayLoad:        "",
		Multimedia:     nil,
		ApnsCollapseId: "",
	}
	stringIos, _ := json.Marshal(iosChannel)

	singleParam := PushSingleParam{
		RequestId: strconv.FormatInt(time.Now().UnixNano(), 10), // 请求唯一标识号
		Audience: &Audience{ // 目标用户
			Cid: []string{
				conf.AndroidCid,
				// conf.IosCid,
			}, // cid推送数组
			Alias:         nil, // 别名送数组
			Tag:           nil, // 推送条件
			FastCustomTag: "",  // 使用用户标签筛选目标用户
		},
		Settings: &Settings{ // 推送条件设置
			TTL: 3600000, // 默认一小时，消息离线时间设置，单位毫秒
			Strategy: &Strategy{ // 厂商通道策略，具体看public_struct.go
				Default: 1,
				Ios:     4,
				St:      1,
				Hw:      1,
				Xm:      1,
				Vv:      1,
				Mz:      1,
				Op:      1,
			},
			Speed:        100, // 推送速度，设置100表示：100条/秒左右，0表示不限速
			ScheduleTime: 0,   // 定时推送时间，必须是7天内的时间，格式：毫秒时间戳
		},
		PushMessage: &PushMessage{
			Duration:     "", // 手机端通知展示时间段
			Notification: nil,
			Transmission: string(stringIos),
			Revoke:       nil,
		},
		PushChannel: &PushChannel{
			Ios: &iosChannel,
			Android: &AndroidChannel{Ups: &Ups{
				Notification: nil,
				TransMission: string(stringIos), // 透传消息内容，与notification 二选一
			}},
		},
	}

	res, err := getui.ToSingleByCid(&singleParam)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(string(res))

	var r *PushSingleResult
	if err := json.Unmarshal(res, &r); err != nil {
		t.Error(err)
		return
	}

	t.Log(r)
	//#endregion
}

func TestConfig_ToSingleByCid2(t *testing.T) {
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

	//#region 具体业务
	iosChannel := IosChannel{
		Type: "",
		Aps: &Aps{
			Alert: &Alert{
				Title: "卡是谁2-" + time.Now().Format(timex.FormatDateTime),
				Body:  "为什么我们每天都要打 TA ？",
			},
			ContentAvailable: 0,
		},
		AutoBadge:      "+1",
		PayLoad:        "",
		Multimedia:     nil,
		ApnsCollapseId: "",
	}
	notification := Notification{
		Title:       "卡是谁3-" + time.Now().Format(timex.FormatDateTime),
		Body:        "为什么我们每天都要打 TA ？",
		ClickType:   "startapp", // 打开应用首页
		BadgeAddNum: 1,
	}

	singleParam := PushSingleParam{
		RequestId: strconv.FormatInt(time.Now().UnixNano(), 10), // 请求唯一标识号
		Audience: &Audience{ // 目标用户
			Cid: []string{
				conf.AndroidCid,
				// conf.IosCid,
			}, // cid推送数组
			Alias:         nil, // 别名送数组
			Tag:           nil, // 推送条件
			FastCustomTag: "",  // 使用用户标签筛选目标用户
		},
		Settings: &Settings{ // 推送条件设置
			TTL: 3600000, // 默认一小时，消息离线时间设置，单位毫秒
			Strategy: &Strategy{ // 厂商通道策略，具体看public_struct.go
				Default: 1,
				Ios:     4,
				St:      4,
				Hw:      4,
				Xm:      4,
				Vv:      4,
				Mz:      4,
				Op:      4,
			},
			Speed:        100, // 推送速度，设置100表示：100条/秒左右，0表示不限速
			ScheduleTime: 0,   // 定时推送时间，必须是7天内的时间，格式：毫秒时间戳
		},
		PushMessage: &PushMessage{
			Duration:     "", // 手机端通知展示时间段
			Notification: &notification,
			Transmission: "",
			Revoke:       nil,
		},
		PushChannel: &PushChannel{
			Ios: &iosChannel,
			Android: &AndroidChannel{Ups: &Ups{
				Notification: &notification,
				TransMission: "", // 透传消息内容，与notification 二选一
			}},
		},
	}
	res, err := getui.ToSingleByCid(&singleParam)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(string(res))

	var r *PushSingleResult
	if err := json.Unmarshal(res, &r); err != nil {
		t.Error(err)
		return
	}

	t.Log(r)
	//#endregion
}
