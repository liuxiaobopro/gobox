package getui

// 推送API文档: https://docs.getui.com/getui/server/rest_v2/push/

// cid单推参数
type PushSingleParam struct {
	RequestId   string       `json:"request_id"`   // 必须字段，请求唯一标识号，10-32位之间；如果request_id重复，会导致消息丢失
	Audience    *Audience    `json:"audience"`     // 必须字段，cid数组，只能填一个cid
	Settings    *Settings    `json:"settings"`     // 非必须，推送条件设置
	PushMessage *PushMessage `json:"push_message"` // 必须字段，个推推送消息参数
	PushChannel *PushChannel `json:"push_channel"` // 非必须，厂商推送消息参数，包含ios消息参数，android厂商消息参数
}

// 群推参数
type PushAppParam struct {
	RequestId   string       `json:"request_id"`   // 必须，请求唯一标识号，10-32位之间；如果request_id重复，会导致消息丢失
	GroupName   string       `json:"group_name"`   // 非必须，任务组名
	Audience    string       `json:"audience"`     // 必须字段，必须为all
	Settings    *Settings    `json:"settings"`     // 非必须，推送条件设置
	PushMessage *PushMessage `json:"push_message"` // 必须字段，个推推送消息参数
	PushChannel *PushChannel `json:"push_channel"` // 非必须，厂商推送消息参数，包含ios消息参数，android厂商消息参数
}

// cid单推返回
type PushSingleResult struct {
	PublicResult
	Data map[string]map[string]string `json:"data"`
}

// 群推返回
type PushAllResult struct {
	PublicResult
	Data map[string]string `json:"data"`
}

// ToSingleByCid 【toSingle】执行cid单推
func (th *Config) ToSingleByCid(param *PushSingleParam) ([]byte, error) {
	if err := th.CheckToken(); err != nil {
		return nil, err
	}

	return th.doPost(PushSingleByCidUrl, param)

}

// ToApp【toApp】执行群推
func (th *Config) ToApp(param *PushAppParam) ([]byte, error) {
	if err := th.CheckToken(); err != nil {
		return nil, err
	}

	return th.doPost(PushAppUrl, param)
}
