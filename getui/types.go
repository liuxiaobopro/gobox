package getui

type AuthReply struct {
	Msg  string     `json:"msg"`
	Code int        `json:"code"`
	Data *authReply `json:"data"`
}

type authReply struct {
	ExpireTime string `json:"expire_time"`
	Token      string `json:"token"`
}
