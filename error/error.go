package error

import replyx "github.com/liuxiaobopro/gobox/reply"

type TOption func(*replyx.T)

func WithCode(code int) TOption {
	return func(t *replyx.T) {
		t.Code = replyx.RespCode(code)
	}
}

func WithData(data interface{}) TOption {
	return func(t *replyx.T) {
		t.Data = data
	}
}

func New(msg string, options ...TOption) error {
	t := &replyx.T{
		Code: replyx.FailErrCode,
		Msg:  msg,
		Data: nil,
	}

	for _, option := range options {
		option(t)
	}

	return t
}

var (
	ErrSucc       = replyx.SuccT          // 操作成功
	ErrFail       = replyx.FailT          // 操作失败
	ErrInternal   = replyx.InternalErrT   // 内部错误
	ErrParam      = replyx.ParamErrT      // 参数错误
	ErrAuth       = replyx.AuthErrT       // 认证失败
	ErrPerm       = replyx.PermErrT       // 无权限
	ErrExist      = replyx.ExistErrT      // 数据已存在
	ErrNotFound   = replyx.NotFoundErrT   // 数据未找到
	ErrLimit      = replyx.LimitErrT      // 已限制操作
	ErrTimeout    = replyx.TimeoutErrT    // 超时错误
	ErrOther      = replyx.OtherErrT      // 未知错误
	ErrMethod     = replyx.MethodErrT     // 方法错误
	ErrTooManyReq = replyx.TooManyReqErrT // 请求频繁
	ErrFileFormat = replyx.FileFormatErrT // 文件格式错误
	ErrEqual      = replyx.EqualErrT      // 与原数据一致
	ErrVersion    = replyx.VersionErrT    // 版本错误
)
