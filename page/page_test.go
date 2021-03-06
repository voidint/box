package page

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPage(t *testing.T) {
	for _, item := range []struct {
		PageNo       int
		PageSize     int
		TotalRecords int
		TotalPages   int
		Offset       int
		Records      []int
	}{
		{PageNo: 1, PageSize: 10, TotalRecords: 100, TotalPages: 10, Offset: 0},
		{PageNo: 2, PageSize: 3, TotalRecords: 7, TotalPages: 3, Offset: 3},
	} {
		pgr := NewPager[int](item.PageNo, item.PageSize, item.TotalRecords)

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
		PageSize     int
		TotalRecords int
		TotalPages   int
		Panic        bool
	}{
		{PageSize: 10, TotalRecords: 10, TotalPages: 1},
		{PageSize: 10, TotalRecords: 0, TotalPages: 0},
		{PageSize: 3, TotalRecords: 10, TotalPages: 4},
		{PageSize: 3, TotalRecords: 0, TotalPages: 0},
		{PageSize: 0, TotalRecords: 10, TotalPages: 0, Panic: true},
		{PageSize: 10, TotalRecords: -1, TotalPages: 0, Panic: true},
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
