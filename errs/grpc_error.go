package errs

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// IsGrpcNotFoundError 判断是否是grpc中的codes.NotFound类型的错误
func IsGrpcNotFoundError(err error) bool {
	return IsGrpcError(err, codes.NotFound)
}

// IsGrpcError 判断是否是grpc中的指定类型错误
func IsGrpcError(err error, code codes.Code) bool {
	if err == nil {
		return false
	}

	if stat, ok := status.FromError(err); ok {
		return stat.Code() == code
	}
	return false
}
