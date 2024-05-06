package errs

import (
	"fmt"
	"net/http"

	"github.com/voidint/box/i18n"
)

// LogicError 业务逻辑错误
type LogicError struct {
	code    int16
	message string
	cause   error
}

// Code 返回错误编码
func (e LogicError) Code() int16 {
	return e.code
}

// Message 返回错误信息
func (e LogicError) Message() string {
	return e.message
}

// Error 实现Error方法
func (e LogicError) Error() string {
	if e.cause == nil {
		return fmt.Sprintf("[%d]%s", e.code, e.message)
	}
	return fmt.Sprintf("[%d]%s\n%+v", e.code, e.message, e.cause)
}

// SetCause 设置根因
func (e *LogicError) SetCause(err error) *LogicError {
	e.cause = err
	return e
}

// GetCause 返回根因错误
func (e *LogicError) GetCause() error {
	return e.cause
}

func WithCause(err error) func(*LogicError) {
	return func(e *LogicError) {
		e.cause = err
	}
}

// NewRawLogicErr 返回指定编码的业务逻辑错误
func NewRawLogicErr(code int16, message string, opts ...func(*LogicError)) *LogicError {
	one := LogicError{
		code:    code,
		message: message,
	}
	for _, opt := range opts {
		opt(&one)
	}
	return &one
}

// NewLogicErr 返回指定编码的业务逻辑错误
func NewLogicErr(code int16, lang string, messageID string, tplData ...any) *LogicError {
	msg, err := i18n.Tr(lang, messageID, tplData...)
	if err != nil {
		msg = messageID
	}
	return NewRawLogicErr(code, msg)
}

// New400LogicError 返回400错误（非法参数）
func New400LogicError(lang string, messageID string, tplData ...any) *LogicError {
	return NewLogicErr(http.StatusBadRequest, lang, messageID, tplData...)
}

// New401LogicError 返回401错误（未认证）
func New401LogicError(lang string, messageID string, tplData ...any) *LogicError {
	return NewLogicErr(http.StatusUnauthorized, lang, messageID, tplData...)
}

// New403LogicError 返回403错误（未授权）
func New403LogicError(lang string, messageID string, tplData ...any) *LogicError {
	return NewLogicErr(http.StatusForbidden, lang, messageID, tplData...)
}

// New404LogicError 返回404错误（资源不存在）
func New404LogicError(lang string, messageID string, tplData ...any) *LogicError {
	return NewLogicErr(http.StatusNotFound, lang, messageID, tplData...)
}

// New500LogicError 返回500错误（服务器内部错误）
func New500LogicError(lang string, messageID string, tplData ...any) *LogicError {
	return NewLogicErr(http.StatusInternalServerError, lang, messageID, tplData...)
}

// IsServerError 返回是否是服务端错误的布尔值
func IsServerError(err error) bool {
	if err == nil {
		return false
	}
	srvErr, ok := err.(*LogicError)
	if !ok {
		return false
	}
	return srvErr.Code() >= http.StatusInternalServerError
}
