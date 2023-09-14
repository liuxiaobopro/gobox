package reply

type RespCode int

type T struct {
	Code RespCode    `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var (
	SuccT          = &T{Code: SuccErrCode, Msg: ErrCodeMsg[SuccErrCode], Data: nil}             // 操作成功
	FailT          = &T{Code: FailErrCode, Msg: ErrCodeMsg[FailErrCode], Data: nil}             // 操作失败
	InternalErrT   = &T{Code: InternalErrCode, Msg: ErrCodeMsg[InternalErrCode], Data: nil}     // 内部错误
	ParamErrT      = &T{Code: ParamErrCode, Msg: ErrCodeMsg[ParamErrCode], Data: nil}           // 参数错误
	AuthErrT       = &T{Code: AuthErrCode, Msg: ErrCodeMsg[AuthErrCode], Data: nil}             // 认证失败
	PermErrT       = &T{Code: PermErrCode, Msg: ErrCodeMsg[PermErrCode], Data: nil}             // 无权限
	ExistErrT      = &T{Code: ExistErrCode, Msg: ErrCodeMsg[ExistErrCode], Data: nil}           // 数据已存在
	NotFoundErrT   = &T{Code: NotFoundErrCode, Msg: ErrCodeMsg[NotFoundErrCode], Data: nil}     // 数据未找到
	LimitErrT      = &T{Code: LimitErrCode, Msg: ErrCodeMsg[LimitErrCode], Data: nil}           // 已限制操作
	TimeoutErrT    = &T{Code: TimeoutErrCode, Msg: ErrCodeMsg[TimeoutErrCode], Data: nil}       // 超时错误
	OtherErrT      = &T{Code: OtherErrCode, Msg: ErrCodeMsg[OtherErrCode], Data: nil}           // 未知错误
	MethodErrT     = &T{Code: MethodErrCode, Msg: ErrCodeMsg[MethodErrCode], Data: nil}         // 方法错误
	TooManyReqErrT = &T{Code: TooManyReqErrCode, Msg: ErrCodeMsg[TooManyReqErrCode], Data: nil} // 请求频繁
	FileFormatErrT = &T{Code: FileFormatErrCode, Msg: ErrCodeMsg[FileFormatErrCode], Data: nil} // 文件格式错误
	EqualErrT      = &T{Code: EqualErrCode, Msg: ErrCodeMsg[EqualErrCode], Data: nil}           // 与原数据一致
)

const (
	SuccErrCode RespCode = iota
	FailErrCode          // 1

	InternalErrCode   // 2
	ParamErrCode      // 3
	AuthErrCode       // 4
	PermErrCode       // 5
	ExistErrCode      // 6
	NotFoundErrCode   // 7
	LimitErrCode      // 8
	TimeoutErrCode    // 9
	OtherErrCode      // 10
	MethodErrCode     // 11
	TooManyReqErrCode // 12
	FileFormatErrCode // 13
	EqualErrCode      // 14
)

var ErrCodeMsg = map[RespCode]string{
	SuccErrCode:       "操作成功",
	FailErrCode:       "操作失败",
	InternalErrCode:   "内部错误",
	ParamErrCode:      "参数错误",
	AuthErrCode:       "认证失败",
	PermErrCode:       "无权限",
	ExistErrCode:      "数据已存在",
	NotFoundErrCode:   "数据未找到",
	LimitErrCode:      "已限制操作",
	TimeoutErrCode:    "超时错误",
	OtherErrCode:      "未知错误",
	MethodErrCode:     "方法错误",
	TooManyReqErrCode: "请求频繁",
	FileFormatErrCode: "文件格式错误",
	EqualErrCode:      "与原数据一致",
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

func (r T) Error() string {
	return r.Msg
}

func Succ(data interface{}) *T {
	return &T{
		Code: SuccErrCode,
		Msg:  ErrCodeMsg[SuccErrCode],
		Data: data,
	}
}

func Fail() *T {
	return &T{
		Code: FailErrCode,
		Msg:  ErrCodeMsg[FailErrCode],
		Data: nil,
	}
}

func InternalErr() *T {
	return &T{
		Code: InternalErrCode,
		Msg:  ErrCodeMsg[InternalErrCode],
		Data: nil,
	}
}

func ParamErr() *T {
	return &T{
		Code: ParamErrCode,
		Msg:  ErrCodeMsg[ParamErrCode],
		Data: nil,
	}
}

func AuthErr() *T {
	return &T{
		Code: AuthErrCode,
		Msg:  ErrCodeMsg[AuthErrCode],
		Data: nil,
	}
}

func PermErr() *T {
	return &T{
		Code: PermErrCode,
		Msg:  ErrCodeMsg[PermErrCode],
		Data: nil,
	}
}

func ExistErr() *T {
	return &T{
		Code: ExistErrCode,
		Msg:  ErrCodeMsg[ExistErrCode],
		Data: nil,
	}
}

func NotFoundErr() *T {
	return &T{
		Code: NotFoundErrCode,
		Msg:  ErrCodeMsg[NotFoundErrCode],
		Data: nil,
	}
}

func LimitErr() *T {
	return &T{
		Code: LimitErrCode,
		Msg:  ErrCodeMsg[LimitErrCode],
		Data: nil,
	}
}

func TimeoutErr() *T {
	return &T{
		Code: TimeoutErrCode,
		Msg:  ErrCodeMsg[TimeoutErrCode],
		Data: nil,
	}
}

func OtherErr() *T {
	return &T{
		Code: OtherErrCode,
		Msg:  ErrCodeMsg[OtherErrCode],
		Data: nil,
	}
}
