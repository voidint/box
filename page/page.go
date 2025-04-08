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
	"github.com/voidint/box/constraints"
	"github.com/voidint/box/db"
)

// Page defines the pagination structure containing metadata and records
type Page[T any, INT constraints.Unsigned] struct {
	PageNo INT `json:"pageNo"` // PageNo specifies the current page number (1-based). For custom JSON field naming,
	// use extra.SetNamingStrategy() from github.com/json-iterator/go.
	PageSize     INT `json:"pageSize"`     // PageSize indicates the number of items per page
	TotalPages   INT `json:"totalPages"`   // TotalPages holds the calculated total number of pages
	TotalRecords INT `json:"totalRecords"` // TotalRecords contains the complete count of dataset records
	Records      []T `json:"records"`      // Records stores the slice of items for the current page
}

// Pager defines the interface for pagination operations
type Pager[T any, INT constraints.Unsigned] interface {
	AddRecords(records ...T)
	BuildDBPage() *db.Page[INT]
	BuildPage() *Page[T, INT]
}

// NewPager constructs a paginator instance with sanitized input parameters
//
// Parameters:
//   - pageNo: current page number (1-based index)
//   - pageSize: number of items per page
//   - totalRecords: total number of records
//
// Returns a properly configured Pager instance
func NewPager[T any, INT constraints.Unsigned](pageNo, pageSize, totalRecords INT) Pager[T, INT] {
	if pageNo <= 0 {
		pageNo = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	if totalRecords < 0 {
		totalRecords = 0
	}
	capacity := pageSize
	if pageSize > totalRecords {
		capacity = totalRecords
	}

	return &pagerImpl[T, INT]{
		pageNo:       pageNo,
		pageSize:     pageSize,
		totalRecords: totalRecords,
		totalPages:   mustCalculateTotalPages(pageSize, totalRecords),
		records:      make([]T, 0, capacity),
	}
}

type pagerImpl[T any, INT constraints.Unsigned] struct {
	pageNo       INT
	pageSize     INT
	totalPages   INT
	totalRecords INT
	records      []T
}

// BuildDBPage constructs database pagination parameters for query execution
func (p *pagerImpl[T, INT]) BuildDBPage() *db.Page[INT] {
	return db.NewPage(p.pageNo, p.pageSize)
}

// AddRecords appends records to the current page's collection
func (p *pagerImpl[T, INT]) AddRecords(records ...T) {
	p.records = append(p.records, records...)
}

// BuildPage assembles and returns the complete pagination structure
func (p *pagerImpl[T, INT]) BuildPage() *Page[T, INT] {
	return &Page[T, INT]{
		PageNo:       p.pageNo,
		PageSize:     p.pageSize,
		TotalPages:   p.totalPages,
		TotalRecords: p.totalRecords,
		Records:      p.records,
	}
}

// EmptyPage initializes a pagination structure with zero values
func EmptyPage[T any, INT constraints.Unsigned](pageNo, pageSize INT) *Page[T, INT] {
	return &Page[T, INT]{
		PageNo:       pageNo,
		PageSize:     pageSize,
		TotalPages:   0,
		TotalRecords: 0,
		Records:      make([]T, 0),
	}
}

func mustCalculateTotalPages[INT constraints.Unsigned](pageSize, totalRecords INT) (totalPages INT) {
	if pageSize <= 0 {
		panic("page size should be positive integer")
	}

	if totalRecords < 0 {
		panic("total records should not be negative integer")
	}

	if totalRecords == 0 {
		return 0
	}

	if totalRecords%pageSize == 0 {
		return totalRecords / pageSize
	}
	return totalRecords/pageSize + 1
}
