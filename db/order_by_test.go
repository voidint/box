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

package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOneOrderBy(t *testing.T) {
	for _, item := range []struct {
		Field     string
		Direction OrderByDirection
		Out1      string
		Out2      []string
	}{
		{
			Field:     "age",
			Direction: DESC,
			Out1:      "age DESC",
			Out2:      []string{"age DESC"},
		},
		{
			Field:     "`age`",
			Direction: ASC,
			Out1:      "`age` ASC",
			Out2:      []string{"`age` ASC"},
		},
		{
			Field:     "",
			Direction: ASC,
			Out1:      " ASC",
			Out2:      []string{" ASC"},
		},
	} {
		assert.Equal(t, item.Out1, OneOrderBy(item.Field, item.Direction).String())
		assert.Equal(t, item.Out2, OneOrderBy(item.Field, item.Direction).Strings())
	}
}

func TestTwoOrderBy(t *testing.T) {
	for _, item := range []struct {
		Field1     string
		Direction1 OrderByDirection
		Field2     string
		Direction2 OrderByDirection
		Out1       string
		Out2       []string
	}{
		{
			Field1:     "age",
			Direction1: DESC,
			Field2:     "gender",
			Direction2: ASC,
			Out1:       "age DESC,gender ASC",
			Out2:       []string{"age DESC", "gender ASC"},
		},
		{
			Field1:     "`age`",
			Direction1: DESC,
			Field2:     "gender",
			Direction2: ASC,
			Out1:       "`age` DESC,gender ASC",
			Out2:       []string{"`age` DESC", "gender ASC"},
		},
		{
			Field1:     "`age`",
			Direction1: ASC,
			Out1:       "`age` ASC, ",
			Out2:       []string{"`age` ASC", " "},
		},
	} {
		assert.Equal(t, item.Out1, TwoOrderBy(item.Field1, item.Direction1, item.Field2, item.Direction2).String())
		assert.Equal(t, item.Out2, TwoOrderBy(item.Field1, item.Direction1, item.Field2, item.Direction2).Strings())
	}
}
