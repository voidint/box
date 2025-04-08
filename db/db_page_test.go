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

func TestDBPage(t *testing.T) {
	for _, item := range []struct {
		InPageNo       uint64
		InPageSize     uint64
		InTotalRecords uint64
		OutLimit       uint64
		OutOffset      uint64
		OutTotalPages  uint64
	}{
		{InPageNo: 1, InPageSize: 10, InTotalRecords: 11, OutLimit: 10, OutOffset: 0, OutTotalPages: 2},
		{InPageNo: 2, InPageSize: 3, InTotalRecords: 10, OutLimit: 3, OutOffset: 3, OutTotalPages: 4},
		{InPageNo: 0, InPageSize: 0, InTotalRecords: 10, OutLimit: 10, OutOffset: 0, OutTotalPages: 1},
		{InPageNo: 1, InPageSize: 30, InTotalRecords: 100, OutLimit: 20, OutOffset: 0, OutTotalPages: 5},
		{InPageNo: 1, InPageSize: 30, InTotalRecords: 22, OutLimit: 20, OutOffset: 0, OutTotalPages: 2},
	} {
		pg := NewPage[uint64](item.InPageNo, item.InPageSize, WithMaxPageSize(uint64(20)))
		assert.Equal(t, item.OutOffset, pg.Offset())
		assert.Equal(t, item.OutLimit, pg.Limit())
		assert.Equal(t, item.OutTotalPages, pg.TotalPages(item.InTotalRecords))
	}
}
