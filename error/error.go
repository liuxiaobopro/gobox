package error

type Code int

func (c Code) Int() int {
	return int(c)
}

type Msg string

func (m Msg) String() string {
	return string(m)
}

type Data interface{}

type T struct {
	Code Code `json:"code"`
	Msg  Msg  `json:"msg"`
	Data Data `json:"data"`
}

type TOption func(*T)

func WithCode(code int) TOption {
	return func(t *T) {
		t.Code = Code(code)
	}
}

func WithData(data interface{}) TOption {
	return func(t *T) {
		t.Data = Data(data)
	}
}

func New(msg string, options ...TOption) error {
	t := &T{
		Code: CodeFail,
		Msg:  Msg(msg),
		Data: nil,
	}

	for _, option := range options {
		option(t)
	}

	return t
}

var (
	ErrSucc       = &T{Code: CodeSucc, Msg: ErrMsg[CodeSucc], Data: nil}             // 操作成功
	ErrFail       = &T{Code: CodeFail, Msg: ErrMsg[CodeFail], Data: nil}             // 操作失败
	ErrInternal   = &T{Code: CodeInternal, Msg: ErrMsg[CodeInternal], Data: nil}     // 内部错误
	ErrParam      = &T{Code: CodeParam, Msg: ErrMsg[CodeParam], Data: nil}           // 参数错误
	ErrAuth       = &T{Code: CodeAuth, Msg: ErrMsg[CodeAuth], Data: nil}             // 认证失败
	ErrPerm       = &T{Code: CodePerm, Msg: ErrMsg[CodePerm], Data: nil}             // 无权限
	ErrExist      = &T{Code: CodeExist, Msg: ErrMsg[CodeExist], Data: nil}           // 数据已存在
	ErrNotFound   = &T{Code: CodeNotFound, Msg: ErrMsg[CodeNotFound], Data: nil}     // 数据未找到
	ErrLimit      = &T{Code: CodeLimit, Msg: ErrMsg[CodeLimit], Data: nil}           // 已限制操作
	ErrTimeout    = &T{Code: CodeTimeout, Msg: ErrMsg[CodeTimeout], Data: nil}       // 超时错误
	ErrOther      = &T{Code: CodeOther, Msg: ErrMsg[CodeOther], Data: nil}           // 未知错误
	ErrMethod     = &T{Code: CodeMethod, Msg: ErrMsg[CodeMethod], Data: nil}         // 方法错误
	ErrTooManyReq = &T{Code: CodeTooManyReq, Msg: ErrMsg[CodeTooManyReq], Data: nil} // 请求频繁
	ErrFileFormat = &T{Code: CodeFileFormat, Msg: ErrMsg[CodeFileFormat], Data: nil} // 文件格式错误
	ErrEqual      = &T{Code: CodeEqual, Msg: ErrMsg[CodeEqual], Data: nil}           // 与原数据一致
	ErrVersion    = &T{Code: CodeVersion, Msg: ErrMsg[CodeVersion], Data: nil}       // 版本错误
)

const (
	CodeSucc Code = iota
	CodeFail
	CodeInternal
	CodeParam
	CodeAuth
	CodePerm
	CodeExist
	CodeNotFound
	CodeLimit
	CodeTimeout
	CodeOther
	CodeMethod
	CodeTooManyReq
	CodeFileFormat
	CodeEqual
	CodeVersion
)

var ErrMsg = map[Code]Msg{
	CodeSucc:       "操作成功",
	CodeFail:       "操作失败",
	CodeInternal:   "内部错误",
	CodeParam:      "参数错误",
	CodeAuth:       "认证失败",
	CodePerm:       "无权限",
	CodeExist:      "数据已存在",
	CodeNotFound:   "数据未找到",
	CodeLimit:      "已限制操作",
	CodeTimeout:    "超时错误",
	CodeOther:      "未知错误",
	CodeMethod:     "方法错误",
	CodeTooManyReq: "请求频繁",
	CodeFileFormat: "文件格式错误",
	CodeEqual:      "与原数据一致",
	CodeVersion:    "版本错误",
}

func (t T) Error() string {
	return t.Msg.String()
}

func (t T) Value() T {
	return t
}
