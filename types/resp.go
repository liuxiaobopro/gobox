package types

type RespCode int

type Resp struct {
	Code RespCode    `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	SuccErrCode RespCode = 0 // 操作成功
	FailErrCode RespCode = 1 // 操作失败

	InternalErrCode RespCode = 1000 // 内部错误
	ParamErrCode    RespCode = 1001 // 参数错误
	AuthErrCode     RespCode = 1002 // 认证错误
	PermErrCode     RespCode = 1003 // 权限错误
	ExistErrCode    RespCode = 1004 // 已存在错误
	NotFoundErrCode RespCode = 1005 // 未找到错误
	LimitErrCode    RespCode = 1006 // 限制错误
	TimeoutErrCode  RespCode = 1007 // 超时错误
	OtherErrCode    RespCode = 1008 // 其他错误
)

func (r Resp) IsSucc() bool {
	return r.Code == SuccErrCode
}

func SuccResp(data interface{}) Resp {
	return Resp{
		Code: SuccErrCode,
		Msg:  "ok",
		Data: data,
	}
}

func FailResp(msg string) Resp {
	return Resp{
		Code: FailErrCode,
		Msg:  msg,
		Data: nil,
	}
}
