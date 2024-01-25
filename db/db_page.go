package db

import (
	"golang.org/x/exp/constraints"
)

// WithMaxPageSize 设置分页大小上限阈值
func WithMaxPageSize[INT constraints.Integer](maxPageSize INT) func(*Page[INT]) {
	return func(pg *Page[INT]) {
		pg.maxPageSize = maxPageSize
	}
}

// Page 分页
type Page[INT constraints.Integer] struct {
	pageNo, pageSize, maxPageSize INT
}

// NewPage 返回分页实例
func NewPage[INT constraints.Integer](pageNo, pageSize INT, opts ...func(*Page[INT])) (pg *Page[INT]) {
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

// Limit 返回分页大小
func (pg *Page[INT]) Limit() INT {
	return pg.pageSize
}

// Offset 返回分页记录偏移量
func (pg *Page[INT]) Offset() INT {
	return (pg.pageNo - 1) * pg.pageSize
}

// PageNo 返回分页页号
func (pg *Page[INT]) PageNo() INT {
	return pg.pageNo
}

// PageSize 返回分页大小
func (pg *Page[INT]) PageSize() INT {
	return pg.pageSize
}

// TotalPages 返回总记录数对应的总分页数
func (pg *Page[INT]) TotalPages(totalRecords INT) (totalPages INT) {
	if totalRecords <= 0 {
		return 0
	}
	if totalRecords%pg.pageSize == 0 {
		return totalRecords / pg.pageSize
	}
	return totalRecords/pg.pageSize + 1
}
