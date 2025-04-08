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

package condexpr

// Any performs conditional selection on generic types
// Returns a if expr evaluates to true, otherwise returns b
func Any[T any](expr bool, a, b T) T {
	if expr {
		return a
	}
	return b
}

// Str performs conditional selection on string type
// Returns a if expr evaluates to true, otherwise returns b
func Str(expr bool, a, b string) string {
	if expr {
		return a
	}
	return b
}

// Int performs conditional selection on int type
// Returns a if expr evaluates to true, otherwise returns b
func Int(expr bool, a, b int) int {
	if expr {
		return a
	}
	return b
}

// Int8 performs conditional selection on int8 type
// Returns a if expr evaluates to true, otherwise returns b
func Int8(expr bool, a, b int8) int8 {
	if expr {
		return a
	}
	return b
}

// Int16 performs conditional selection on int16 type
// Returns a if expr evaluates to true, otherwise returns b
func Int16(expr bool, a, b int16) int16 {
	if expr {
		return a
	}
	return b
}

// Int32 performs conditional selection on int32 type
// Returns a if expr evaluates to true, otherwise returns b
func Int32(expr bool, a, b int32) int32 {
	if expr {
		return a
	}
	return b
}

// Int64 performs conditional selection on int64 type
// Returns a if expr evaluates to true, otherwise returns b
func Int64(expr bool, a, b int64) int64 {
	if expr {
		return a
	}
	return b
}

// Uint performs conditional selection on uint type
// Returns a if expr evaluates to true, otherwise returns b
func Uint(expr bool, a, b uint) uint {
	if expr {
		return a
	}
	return b
}

// Uint8 performs conditional selection on uint8 type
// Returns a if expr evaluates to true, otherwise returns b
func Uint8(expr bool, a, b uint8) uint8 {
	if expr {
		return a
	}
	return b
}

// Uint16 performs conditional selection on uint16 type
// Returns a if expr evaluates to true, otherwise returns b
func Uint16(expr bool, a, b uint16) uint16 {
	if expr {
		return a
	}
	return b
}

// Uint32 performs conditional selection on uint32 type
// Returns a if expr evaluates to true, otherwise returns b
func Uint32(expr bool, a, b uint32) uint32 {
	if expr {
		return a
	}
	return b
}

// Uint64 performs conditional selection on uint64 type
// Returns a if expr evaluates to true, otherwise returns b
func Uint64(expr bool, a, b uint64) uint64 {
	if expr {
		return a
	}
	return b
}

// Float32 performs conditional selection on float32 type
// Returns a if expr evaluates to true, otherwise returns b
func Float32(expr bool, a, b float32) float32 {
	if expr {
		return a
	}
	return b
}

// Float64 performs conditional selection on float64 type
// Returns a if expr evaluates to true, otherwise returns b
func Float64(expr bool, a, b float64) float64 {
	if expr {
		return a
	}
	return b
}
