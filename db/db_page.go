package db

import (
	"golang.org/x/exp/constraints"
)

// Page 分页
type Page[INT constraints.Integer] struct {
	pageNo, pageSize INT
}

// NewPage 返回分页实例
func NewPage[INT constraints.Integer](pageNo, pageSize INT) (pg *Page[INT]) {
	if pageNo <= 0 {
		pageNo = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	return &Page[INT]{
		pageNo:   pageNo,
		pageSize: pageSize,
	}
}

// Limit 返回分页大小
func (pg *Page[INT]) Limit() INT {
	return pg.pageSize
}

// Offset 返回分页记录偏移量
func (pg *Page[INT]) Offset() INT {
	return (pg.pageNo - 1) * pg.pageSize
}

// 返回分页页号
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
