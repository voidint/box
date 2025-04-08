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
	"bytes"
	"fmt"
)

// OrderByDirection defines sorting direction for SQL ORDER BY clause
type OrderByDirection string

var (
	// ASC ascending sort order
	ASC OrderByDirection = "ASC"
	// DESC descending sort order
	DESC OrderByDirection = "DESC"
)

// OrderByPair represents a field and its sorting direction pair
type OrderByPair struct {
	Name      string
	Direction OrderByDirection
}

func (pair OrderByPair) String() string {
	return fmt.Sprintf("%s %s", pair.Name, string(pair.Direction))
}

// OrderBy collects multiple sorting criteria for query
type OrderBy []OrderByPair

// String formats sorting criteria as SQL ORDER BY clause
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

// Strings returns slice of formatted sorting expressions
func (ob OrderBy) Strings() []string {
	if len(ob) <= 0 {
		return nil
	}

	items := make([]string, 0, len(ob))
	for i := range ob {
		items = append(items, ob[i].String())
	}
	return items
}

// OneOrderBy creates OrderBy with single sorting pair
func OneOrderBy(name string, direction OrderByDirection) OrderBy {
	return OrderBy([]OrderByPair{
		{
			Name:      name,
			Direction: direction,
		},
	})
}

// TwoOrderBy creates OrderBy with two sorting pairs
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
