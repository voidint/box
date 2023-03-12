package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractFromPtr(t *testing.T) {
	s1 := ""
	s2 := "box"

	type args struct {
		p *string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "入参是nil",
			args: args{p: nil},
			want: "",
		},
		{
			name: "入参是空字符串",
			args: args{p: &s1},
			want: s1,
		},
		{
			name: "入参是非空字符串",
			args: args{p: &s2},
			want: s2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ExtractFromPtr(tt.args.p))
		})
	}
}
