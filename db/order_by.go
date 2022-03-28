package db

import (
	"bytes"
	"fmt"
)

// OrderByDirection SQL(ORDER BY)排序方向
type OrderByDirection string

var (
	// ASC 升序
	ASC OrderByDirection = "ASC"
	// DESC 降序
	DESC OrderByDirection = "DESC"
)

// OrderByPair ORDER BY对
type OrderByPair struct {
	Name      string
	Direction OrderByDirection
}

func (pair OrderByPair) String() string {
	return fmt.Sprintf("%s %s", pair.Name, string(pair.Direction))
}

// OrderBy ORDER BY信息
type OrderBy []OrderByPair

func (ob OrderBy) String() string {
	if len(ob) <= 0 {
		return ""
	}

	var buf bytes.Buffer
	for i := range ob {
		buf.WriteString(ob[i].String())
		if i < len(ob)-1 {
			buf.WriteByte(',')
		}
	}
	return buf.String()
}

// OneOrderBy 构建仅包含一对的orderBy
func OneOrderBy(name string, direction OrderByDirection) OrderBy {
	return OrderBy([]OrderByPair{
		{
			Name:      name,
			Direction: direction,
		},
	})
}

// TwoOrderBy 构建包含两对的orderBy
func TwoOrderBy(name0 string, direction0 OrderByDirection, name1 string, direction1 OrderByDirection) OrderBy {
	return OrderBy([]OrderByPair{
		{
			Name:      name0,
			Direction: direction0,
		},
		{
			Name:      name1,
			Direction: direction1,
		},
	})
}
