package resp

type RespCode int

type T struct {
	Code RespCode    `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Pt *T

var (
	SuccT        = &T{Code: SuccErrCode, Msg: ErrCodeMsg[SuccErrCode], Data: nil}
	FailT        = &T{Code: FailErrCode, Msg: ErrCodeMsg[FailErrCode], Data: nil}
	InternalErrT = &T{Code: InternalErrCode, Msg: ErrCodeMsg[InternalErrCode], Data: nil}
	ParamErrT    = &T{Code: ParamErrCode, Msg: ErrCodeMsg[ParamErrCode], Data: nil}
	AuthErrT     = &T{Code: AuthErrCode, Msg: ErrCodeMsg[AuthErrCode], Data: nil}
	PermErrT     = &T{Code: PermErrCode, Msg: ErrCodeMsg[PermErrCode], Data: nil}
	ExistErrT    = &T{Code: ExistErrCode, Msg: ErrCodeMsg[ExistErrCode], Data: nil}
	NotFoundErrT = &T{Code: NotFoundErrCode, Msg: ErrCodeMsg[NotFoundErrCode], Data: nil}
	LimitErrT    = &T{Code: LimitErrCode, Msg: ErrCodeMsg[LimitErrCode], Data: nil}
	TimeoutErrT  = &T{Code: TimeoutErrCode, Msg: ErrCodeMsg[TimeoutErrCode], Data: nil}
	OtherErrT    = &T{Code: OtherErrCode, Msg: ErrCodeMsg[OtherErrCode], Data: nil}
)

const (
	SuccErrCode RespCode = iota
	FailErrCode          // 1

	InternalErrCode // 2
	ParamErrCode    // 3
	AuthErrCode     // 4
	PermErrCode     // 5
	ExistErrCode    // 6
	NotFoundErrCode // 7
	LimitErrCode    // 8
	TimeoutErrCode  // 9
	OtherErrCode    // 10
)

var ErrCodeMsg = map[RespCode]string{
	SuccErrCode:     "ok",
	FailErrCode:     "操作失败",
	InternalErrCode: "内部错误",
	ParamErrCode:    "参数错误",
	AuthErrCode:     "认证失败",
	PermErrCode:     "无权限",
	ExistErrCode:    "数据已存在",
	NotFoundErrCode: "数据未找到",
	LimitErrCode:    "已限制操作",
	TimeoutErrCode:  "超时错误",
	OtherErrCode:    "未知错误",
}

func (r T) IsSucc() bool {
	return r.Code == SuccErrCode
}

func (r T) GetCode() RespCode {
	return r.Code
}

func (r T) GetMsg() string {
	return r.Msg
}

func (r T) GetData() interface{} {
	return r.Data
}

func (r T) ToPt() *Pt {
	Pt := Pt(&r)
	return &Pt
}

func Succ(data interface{}) Pt {
	return &T{
		Code: SuccErrCode,
		Msg:  ErrCodeMsg[SuccErrCode],
		Data: data,
	}
}

func Fail() Pt {
	return &T{
		Code: FailErrCode,
		Msg:  ErrCodeMsg[FailErrCode],
		Data: nil,
	}
}

func InternalErr() Pt {
	return &T{
		Code: InternalErrCode,
		Msg:  ErrCodeMsg[InternalErrCode],
		Data: nil,
	}
}

func ParamErr() Pt {
	return &T{
		Code: ParamErrCode,
		Msg:  ErrCodeMsg[ParamErrCode],
		Data: nil,
	}
}

func AuthErr() Pt {
	return &T{
		Code: AuthErrCode,
		Msg:  ErrCodeMsg[AuthErrCode],
		Data: nil,
	}
}

func PermErr() Pt {
	return &T{
		Code: PermErrCode,
		Msg:  ErrCodeMsg[PermErrCode],
		Data: nil,
	}
}

func ExistErr() Pt {
	return &T{
		Code: ExistErrCode,
		Msg:  ErrCodeMsg[ExistErrCode],
		Data: nil,
	}
}

func NotFoundErr() Pt {
	return &T{
		Code: NotFoundErrCode,
		Msg:  ErrCodeMsg[NotFoundErrCode],
		Data: nil,
	}
}

func LimitErr() Pt {
	return &T{
		Code: LimitErrCode,
		Msg:  ErrCodeMsg[LimitErrCode],
		Data: nil,
	}
}

func TimeoutErr() Pt {
	return &T{
		Code: TimeoutErrCode,
		Msg:  ErrCodeMsg[TimeoutErrCode],
		Data: nil,
	}
}

func OtherErr() Pt {
	return &T{
		Code: OtherErrCode,
		Msg:  ErrCodeMsg[OtherErrCode],
		Data: nil,
	}
}
