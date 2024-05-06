package error

// import "github.com/liuxiaobopro/gobox/reply"

// type Code int

// type Err struct {
// 	Code Code        `json:"code"`
// 	Msg  string      `json:"msg"`
// 	Data interface{} `json:"data"`
// }

// // var (
// // 	SuccT          = &T{Code: SuccErrCode, Msg: ErrCodeMsg[SuccErrCode], Data: nil}             // 操作成功
// // 	FailT          = &T{Code: FailErrCode, Msg: ErrCodeMsg[FailErrCode], Data: nil}             // 操作失败
// // 	InternalErrT   = &T{Code: InternalErrCode, Msg: ErrCodeMsg[InternalErrCode], Data: nil}     // 内部错误
// // 	ParamErrT      = &T{Code: ParamErrCode, Msg: ErrCodeMsg[ParamErrCode], Data: nil}           // 参数错误
// // 	AuthErrT       = &T{Code: AuthErrCode, Msg: ErrCodeMsg[AuthErrCode], Data: nil}             // 认证失败
// // 	PermErrT       = &T{Code: PermErrCode, Msg: ErrCodeMsg[PermErrCode], Data: nil}             // 无权限
// // 	ExistErrT      = &T{Code: ExistErrCode, Msg: ErrCodeMsg[ExistErrCode], Data: nil}           // 数据已存在
// // 	NotFoundErrT   = &T{Code: NotFoundErrCode, Msg: ErrCodeMsg[NotFoundErrCode], Data: nil}     // 数据未找到
// // 	LimitErrT      = &T{Code: LimitErrCode, Msg: ErrCodeMsg[LimitErrCode], Data: nil}           // 已限制操作
// // 	TimeoutErrT    = &T{Code: TimeoutErrCode, Msg: ErrCodeMsg[TimeoutErrCode], Data: nil}       // 超时错误
// // 	OtherErrT      = &T{Code: OtherErrCode, Msg: ErrCodeMsg[OtherErrCode], Data: nil}           // 未知错误
// // 	MethodErrT     = &T{Code: MethodErrCode, Msg: ErrCodeMsg[MethodErrCode], Data: nil}         // 方法错误
// // 	TooManyReqErrT = &T{Code: TooManyReqErrCode, Msg: ErrCodeMsg[TooManyReqErrCode], Data: nil} // 请求频繁
// // 	FileFormatErrT = &T{Code: FileFormatErrCode, Msg: ErrCodeMsg[FileFormatErrCode], Data: nil} // 文件格式错误
// // 	EqualErrT      = &T{Code: EqualErrCode, Msg: ErrCodeMsg[EqualErrCode], Data: nil}           // 与原数据一致
// // 	VersionErrT    = &T{Code: VersionErrCode, Msg: ErrCodeMsg[VersionErrCode], Data: nil}       // 版本错误
// // )

// const (
// 	SuccErrCode Code = iota
// 	FailErrCode          // 1

// 	InternalErrCode   // 2
// 	ParamErrCode      // 3
// 	AuthErrCode       // 4
// 	PermErrCode       // 5
// 	ExistErrCode      // 6
// 	NotFoundErrCode   // 7
// 	LimitErrCode      // 8
// 	TimeoutErrCode    // 9
// 	OtherErrCode      // 10
// 	MethodErrCode     // 11
// 	TooManyReqErrCode // 12
// 	FileFormatErrCode // 13
// 	EqualErrCode      // 14
// 	VersionErrCode    // 15
// )

// var ErrCodeMsg = map[Code]string{
// 	SuccErrCode:       "操作成功",
// 	FailErrCode:       "操作失败",
// 	InternalErrCode:   "内部错误",
// 	ParamErrCode:      "参数错误",
// 	AuthErrCode:       "认证失败",
// 	PermErrCode:       "无权限",
// 	ExistErrCode:      "数据已存在",
// 	NotFoundErrCode:   "数据未找到",
// 	LimitErrCode:      "已限制操作",
// 	TimeoutErrCode:    "超时错误",
// 	OtherErrCode:      "未知错误",
// 	MethodErrCode:     "方法错误",
// 	TooManyReqErrCode: "请求频繁",
// 	FileFormatErrCode: "文件格式错误",
// 	EqualErrCode:      "与原数据一致",
// 	VersionErrCode:    "版本错误",
// }

// var (
// 	Succ = &Err{Code: SuccCode, Msg: ErrMsg[SuccCode], Data: nil}             // 操作成功
// )
