package null

import (
	"math"
	"testing"

	"github.com/guregu/null/v5"
	"github.com/stretchr/testify/assert"
)

func TestMustIntFromUint64Ptr(t *testing.T) {
	var v1 uint64 = 0
	var v2 uint64 = 7
	var v3 uint64 = math.MaxInt64
	var v4 uint64 = math.MaxInt64 + 1

	type args struct {
		v *uint64
	}
	tests := []struct {
		name  string
		args  args
		want  null.Int
		panic bool
	}{
		{
			name: "入参为nil",
			args: args{v: nil},
			want: null.NewInt(0, false),
		},
		{
			name: "入参为0",
			args: args{v: &v1},
			want: null.NewInt(int64(v1), true),
		},
		{
			name: "入参为正整数",
			args: args{v: &v2},
			want: null.NewInt(int64(v2), true),
		},
		{
			name: "入参为math.MaxInt64",
			args: args{v: &v3},
			want: null.NewInt(int64(v3), true),
		},
		{
			name:  "入参为极大值（math.MaxInt64+1）",
			args:  args{v: &v4},
			panic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panic {
				assert.Panics(t, func() { MustIntFromUint64Ptr(tt.args.v) })

			} else {
				assert.Equal(t, tt.want, MustIntFromUint64Ptr(tt.args.v))
			}

		})
	}
}

func TestMustIntFromUint32Ptr(t *testing.T) {
	var v1 uint32 = 0
	var v2 uint32 = 7
	var v3 uint32 = math.MaxInt32
	var v4 uint32 = math.MaxInt32 + 1

	type args struct {
		v *uint32
	}
	tests := []struct {
		name  string
		args  args
		want  null.Int
		panic bool
	}{
		{
			name: "入参为nil",
			args: args{v: nil},
			want: null.NewInt(0, false),
		},
		{
			name: "入参为0",
			args: args{v: &v1},
			want: null.NewInt(int64(v1), true),
		},
		{
			name: "入参为正整数",
			args: args{v: &v2},
			want: null.NewInt(int64(v2), true),
		},
		{
			name: "入参为math.MaxInt32",
			args: args{v: &v3},
			want: null.NewInt(int64(v3), true),
		},
		{
			name:  "入参为极大值（math.MaxInt32+1）",
			args:  args{v: &v4},
			panic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panic {
				assert.Panics(t, func() { MustIntFromUint32Ptr(tt.args.v) })

			} else {
				assert.Equal(t, tt.want, MustIntFromUint32Ptr(tt.args.v))
			}
		})
	}
}
