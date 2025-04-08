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

package numeric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractFromPtr(t *testing.T) {
	v1 := uint8(1)
	v2 := uint16(2)
	v3 := uint32(3)
	v4 := uint64(4)
	v5 := uint(5)
	v6 := float32(6)
	v7 := float64(7)

	assert.Equal(t, uint8(0), ExtractFromPtr[uint8](nil))
	assert.Equal(t, v1, ExtractFromPtr(&v1))
	assert.Equal(t, v2, ExtractFromPtr(&v2))
	assert.Equal(t, v3, ExtractFromPtr(&v3))
	assert.Equal(t, v4, ExtractFromPtr(&v4))
	assert.Equal(t, v5, ExtractFromPtr(&v5))
	assert.Equal(t, v6, ExtractFromPtr(&v6))
	assert.Equal(t, v7, ExtractFromPtr(&v7))
}

func TestJoin(t *testing.T) {
	assert.Equal(t, "", Join([]uint64{}, "|"))
	assert.Equal(t, "1", Join([]uint64{1}, "|"))
	assert.Equal(t, "1|2", Join([]uint64{1, 2}, "|"))
	assert.Equal(t, "1|2|3", Join([]uint64{1, 2, 3}, "|"))
	assert.Equal(t, "1|2|3|4", Join([]uint64{1, 2, 3, 4}, "|"))
	assert.Equal(t, "1234", Join([]uint64{1, 2, 3, 4}, ""))
	assert.Equal(t, "1 2 3 4", Join([]uint64{1, 2, 3, 4}, " "))
	assert.Equal(t, "1|2|3|4", Join([]float64{1, 2, 3, 4}, "|"))
	assert.Equal(t, "1.12|2.23|3.34|4.45", Join([]float64{1.12, 2.23, 3.34, 4.45}, "|"))
}
