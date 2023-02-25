package errs

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// IsGrpcNotFoundError 判断是否是grpc中的codes.NotFound类型的错误
func IsGrpcNotFoundError(err error) bool {
	if err == nil {
		return false
	}
	if stat, ok := status.FromError(err); ok {
		switch stat.Code() {
		case codes.NotFound:
			return true
		}
	}
	return false
}
