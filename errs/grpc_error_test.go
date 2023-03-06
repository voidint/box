package errs

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestIsGrpcNotFoundError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "入参是nil",
			args: args{err: nil},
			want: false,
		},
		{
			name: "入参是普通error",
			args: args{err: errors.New("invalid arguments")},
			want: false,
		},
		{
			name: "入参是grpc中的codes.Internal类型错误",
			args: args{err: status.Error(codes.Internal, "internal error")},
			want: false,
		},
		{
			name: "入参是grpc中的codes.NotFound类型错误",
			args: args{err: status.Error(codes.NotFound, "not found error")},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, IsGrpcNotFoundError(tt.args.err))
		})
	}
}

func TestIsGrpcError(t *testing.T) {
	type args struct {
		err  error
		code codes.Code
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "入参是nil",
			args: args{err: nil, code: codes.Internal},
			want: false,
		},
		{
			name: "入参是普通error",
			args: args{err: errors.New("invalid arguments"), code: codes.Internal},
			want: false,
		},
		{
			name: "入参是grpc中的codes.Internal类型错误",
			args: args{err: status.Error(codes.Internal, "internal error"), code: codes.Internal},
			want: true,
		},
		{
			name: "入参是grpc中的codes.NotFound类型错误",
			args: args{err: status.Error(codes.NotFound, "not found error"), code: codes.Internal},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, IsGrpcError(tt.args.err, tt.args.code))
		})
	}
}
