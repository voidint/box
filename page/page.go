package page

import (
	"github.com/voidint/box/db"
	"golang.org/x/exp/constraints"
)

// Integer 整型类型约束
type Integer interface {
	constraints.Integer
}

// Page 分页结构定义
type Page[T any, INT Integer] struct {
	PageNo       INT `json:"pageNo"` // 若序列化时想要使用其他的命名风格，建议使用 github.com/json-iterator/go 库中的 extra.SetNamingStrategy() 函数注册自定义命名策略。
	PageSize     INT `json:"pageSize"`
	TotalPages   INT `json:"totalPages"`
	TotalRecords INT `json:"totalRecords"`
	Records      []T `json:"records"`
}

// Pager 分页接口
type Pager[T any, INT Integer] interface {
	AddRecords(records ...T)
	BuildDBPage() *db.Page[INT]
	BuildPage() *Page[T, INT]
}

// NewPager 构建一个分页对象
func NewPager[T any, INT Integer](pageNo, pageSize, totalRecords INT) Pager[T, INT] {
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

// pagerImpl 实际的分页对象
type pagerImpl[T any, INT Integer] struct {
	pageNo       INT
	pageSize     INT
	totalPages   INT
	totalRecords INT
	records      []T
}

// BuildDBPage 返回数据库分页
func (p *pagerImpl[T, INT]) BuildDBPage() *db.Page[INT] {
	return db.NewPage(p.pageNo, p.pageSize)
}

// AddRecords 往分页中追加记录
func (p *pagerImpl[T, INT]) AddRecords(records ...T) {
	p.records = append(p.records, records...)
}

// BuildPage 构造一个分页对象并返回
func (p *pagerImpl[T, INT]) BuildPage() *Page[T, INT] {
	return &Page[T, INT]{
		PageNo:       p.pageNo,
		PageSize:     p.pageSize,
		TotalPages:   p.totalPages,
		TotalRecords: p.totalRecords,
		Records:      p.records,
	}
}

// EmptyPage 返回空分页
func EmptyPage[T any, INT Integer](pageNo, pageSize INT) *Page[T, INT] {
	return &Page[T, INT]{
		PageNo:       pageNo,
		PageSize:     pageSize,
		TotalPages:   0,
		TotalRecords: 0,
		Records:      make([]T, 0),
	}
}

// mustCalculateTotalPages 计算总分页数
func mustCalculateTotalPages[INT Integer](pageSize, totalRecords INT) (totalPages INT) {
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
