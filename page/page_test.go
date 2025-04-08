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

package page

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPage(t *testing.T) {
	for _, item := range []struct {
		PageNo       uint64
		PageSize     uint64
		TotalRecords uint64
		TotalPages   uint64
		Offset       uint64
		Records      []uint64
	}{
		{PageNo: 1, PageSize: 10, TotalRecords: 100, TotalPages: 10, Offset: 0},
		{PageNo: 2, PageSize: 3, TotalRecords: 7, TotalPages: 3, Offset: 3},
	} {
		pgr := NewPager[uint64](item.PageNo, item.PageSize, item.TotalRecords)

		pg := pgr.BuildPage()
		assert.NotNil(t, pg)
		assert.Equal(t, item.TotalPages, pg.TotalPages)

		dbPg := pgr.BuildDBPage()
		assert.NotNil(t, dbPg)
		assert.Equal(t, item.PageSize, dbPg.Limit())
		assert.Equal(t, item.Offset, dbPg.Offset())
		assert.Equal(t, item.PageNo, dbPg.PageNo())
		assert.Equal(t, item.PageSize, dbPg.PageSize())
	}
}

func TestMustCalculateTotalPages(t *testing.T) {
	for _, item := range []struct {
		PageSize     uint64
		TotalRecords uint64
		TotalPages   uint64
		Panic        bool
	}{
		{PageSize: 10, TotalRecords: 10, TotalPages: 1},
		{PageSize: 10, TotalRecords: 0, TotalPages: 0},
		{PageSize: 3, TotalRecords: 10, TotalPages: 4},
		{PageSize: 3, TotalRecords: 0, TotalPages: 0},
		{PageSize: 0, TotalRecords: 10, TotalPages: 0, Panic: true},
	} {
		if item.Panic {
			assert.Panics(t, func() {
				mustCalculateTotalPages(item.PageSize, item.TotalRecords)
			})
		} else {
			mustCalculateTotalPages(item.PageSize, item.TotalRecords)
		}
	}
}
