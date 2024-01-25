package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDBPage(t *testing.T) {
	for _, item := range []struct {
		InPageNo       int
		InPageSize     int
		InTotalRecords int
		OutLimit       int
		OutOffset      int
		OutTotalPages  int
	}{
		{InPageNo: 1, InPageSize: 10, InTotalRecords: 11, OutLimit: 10, OutOffset: 0, OutTotalPages: 2},
		{InPageNo: 2, InPageSize: 3, InTotalRecords: 10, OutLimit: 3, OutOffset: 3, OutTotalPages: 4},
		{InPageNo: 0, InPageSize: 0, InTotalRecords: 10, OutLimit: 10, OutOffset: 0, OutTotalPages: 1},
		{InPageNo: -1, InPageSize: -2, InTotalRecords: 0, OutLimit: 10, OutOffset: 0, OutTotalPages: 0},
		{InPageNo: 1, InPageSize: 30, InTotalRecords: 100, OutLimit: 20, OutOffset: 0, OutTotalPages: 5},
		{InPageNo: 1, InPageSize: 30, InTotalRecords: 22, OutLimit: 20, OutOffset: 0, OutTotalPages: 2},
	} {
		pg := NewPage(item.InPageNo, item.InPageSize, WithMaxPageSize(20))
		assert.Equal(t, item.OutOffset, pg.Offset())
		assert.Equal(t, item.OutLimit, pg.Limit())
		assert.Equal(t, item.OutTotalPages, pg.TotalPages(item.InTotalRecords))
	}
}
