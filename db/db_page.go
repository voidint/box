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
	"github.com/voidint/box/constraints"
)

// WithMaxPageSize configures the maximum allowed page size
// for pagination parameters validation.
func WithMaxPageSize[INT constraints.Unsigned](maxPageSize INT) func(*Page[INT]) {
	return func(pg *Page[INT]) {
		pg.maxPageSize = maxPageSize
	}
}

// Page defines pagination parameters for query operations
type Page[INT constraints.Unsigned] struct {
	pageNo      INT // Current page number (1-based)
	pageSize    INT // Number of items per page
	maxPageSize INT // Maximum allowed page size (0 means no limit)
}

// NewPage creates pagination parameters with validation
func NewPage[INT constraints.Unsigned](pageNo, pageSize INT, opts ...func(*Page[INT])) (pg *Page[INT]) {
	pg = &Page[INT]{
		pageNo:   pageNo,
		pageSize: pageSize,
	}

	if pageNo <= 0 {
		pg.pageNo = 1
	}

	if pageSize <= 0 {
		pg.pageSize = 10
	}

	for _, setter := range opts {
		setter(pg)
	}

	if pg.maxPageSize > 0 && pg.pageSize > pg.maxPageSize {
		pg.pageSize = pg.maxPageSize
	}

	return
}

// Limit returns configured items per page
func (pg *Page[INT]) Limit() INT {
	return pg.pageSize
}

// Offset returns SQL OFFSET clause value
func (pg *Page[INT]) Offset() INT {
	return (pg.pageNo - 1) * pg.pageSize
}

// PageNo returns current page number
func (pg *Page[INT]) PageNo() INT {
	return pg.pageNo
}

// PageSize returns configured items per page
func (pg *Page[INT]) PageSize() INT {
	return pg.pageSize
}

// TotalPages calculates total pages based on total records count
func (pg *Page[INT]) TotalPages(totalRecords INT) (totalPages INT) {
	if totalRecords <= 0 {
		return 0
	}
	if totalRecords%pg.pageSize == 0 {
		return totalRecords / pg.pageSize
	}
	return totalRecords/pg.pageSize + 1
}
