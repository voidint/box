// Copyright (c) 2025 voidint <voidint@126.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package errs

import (
	"fmt"
	"net/http"

	"github.com/voidint/box/i18n"
)

// LogicError represents a business logic error with code mapping and localization support
type LogicError struct {
	code    int16
	message string
	cause   error
}

// Code returns the HTTP status code mapped to this business error
func (e LogicError) Code() int16 {
	return e.code
}

// Message returns the localized error message for client presentation
func (e LogicError) Message() string {
	return e.message
}

// Error implements error interface with formatted error details
func (e LogicError) Error() string {
	if e.cause == nil {
		return fmt.Sprintf("[%d]%s", e.code, e.message)
	}
	return fmt.Sprintf("[%d]%s\n%+v", e.code, e.message, e.cause)
}

// SetCause attaches the root cause error for debugging purposes
func (e *LogicError) SetCause(err error) *LogicError {
	e.cause = err
	return e
}

// GetCause retrieves the underlying error that triggered this business error
func (e *LogicError) GetCause() error {
	return e.cause
}

func WithCause(err error) func(*LogicError) {
	return func(e *LogicError) {
		e.cause = err
	}
}

// NewRawLogicErr creates a business error with raw message (non-localized)
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

// NewLogicErr creates a localized business error using i18n message ID
func NewLogicErr(code int16, lang string, messageID string, tplData ...any) *LogicError {
	msg, err := i18n.Tr(lang, messageID, tplData...)
	if err != nil {
		msg = messageID
	}
	return NewRawLogicErr(code, msg)
}

// New400LogicError creates Bad Request error (HTTP 400)
// Typically indicates invalid request parameters or malformed input
func New400LogicError(lang string, messageID string, tplData ...any) *LogicError {
	return NewLogicErr(http.StatusBadRequest, lang, messageID, tplData...)
}

// New401LogicError creates Unauthorized error (HTTP 401)
// Indicates missing or invalid authentication credentials
func New401LogicError(lang string, messageID string, tplData ...any) *LogicError {
	return NewLogicErr(http.StatusUnauthorized, lang, messageID, tplData...)
}

// New403LogicError creates Forbidden error (HTTP 403)
// The client lacks sufficient permissions for the requested operation
func New403LogicError(lang string, messageID string, tplData ...any) *LogicError {
	return NewLogicErr(http.StatusForbidden, lang, messageID, tplData...)
}

// New404LogicError creates Not Found error (HTTP 404)
// Specified resource does not exist in the system
func New404LogicError(lang string, messageID string, tplData ...any) *LogicError {
	return NewLogicErr(http.StatusNotFound, lang, messageID, tplData...)
}

// New500LogicError creates Internal Server Error (HTTP 500)
// Reserved for unexpected server-side failures
func New500LogicError(lang string, messageID string, tplData ...any) *LogicError {
	return NewLogicErr(http.StatusInternalServerError, lang, messageID, tplData...)
}

// IsServerError checks if the error represents a server-side issue (HTTP 5xx)
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
