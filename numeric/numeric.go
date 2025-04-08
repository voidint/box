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
	"fmt"
	"strings"
)

// PositiveToPtr returns a pointer to v if it's greater than zero, otherwise returns nil.
func PositiveToPtr[T ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint | ~float32 | ~float64](v T) *T {
	if v > 0 {
		return &v
	}
	return nil
}

// PositiveUint8ToUint32Ptr converts uint8 to uint32 pointer with null-safety.
func PositiveUint8ToUint32Ptr(v uint8) *uint32 {
	if v == 0 {
		return nil
	}
	ret := uint32(v)
	return &ret
}

// ExtractFromPtr provides null-safe value extraction from pointer.
func ExtractFromPtr[T ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint | ~float32 | ~float64](p *T) T {
	if p == nil {
		return 0
	}
	return *p
}

type Number interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint | ~int8 | ~int16 | ~int32 | ~int64 | ~int | ~float32 | ~float64
}

// Join concatenates numeric values into a string with specified separator.
func Join[T Number](items []T, sep string) string {
	switch len(items) {
	case 0:
		return ""
	case 1:
		return fmt.Sprintf("%v", items[0])
	case 2:
		return fmt.Sprintf("%v%s%v", items[0], sep, items[1])
	case 3:
		return fmt.Sprintf("%v%s%v%s%v", items[0], sep, items[1], sep, items[2])
	default:
		var buf strings.Builder
		for i := range items {
			buf.WriteString(fmt.Sprintf("%v", items[i]))
			if i != len(items)-1 {
				buf.WriteString(sep)
			}
		}
		return buf.String()
	}
}
