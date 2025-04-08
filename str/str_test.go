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

func TestNonEmptyToPtr(t *testing.T) {
	s1 := " 	"
	s2 := "box"

	type args struct {
		v string
	}
	tests := []struct {
		name string
		args args
		want *string
	}{
		{
			name: "入参是空字符串",
			args: args{v: ""},
			want: nil,
		},
		{
			name: "入参是空白字符串",
			args: args{v: s1},
			want: &s1,
		},
		{
			name: "入参是非空字符串",
			args: args{v: s2},
			want: &s2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NonEmptyToPtr(tt.args.v))
		})
	}
}
