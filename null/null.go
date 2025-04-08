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

package null

import (
	"math"
	"time"

	"github.com/guregu/null/v5"
)

// MustIntFromUint64Ptr performs null-safe conversion from uint64 pointer to null.Int.
func MustIntFromUint64Ptr(v *uint64) null.Int {
	if v == nil {
		return null.NewInt(0, false)
	}
	if *v > math.MaxInt64 {
		panic("parameter value too large")
	}
	return null.IntFrom(int64(*v))
}

// MustIntFromUint32Ptr converts uint32 pointer to null.Int with type checking.
func MustIntFromUint32Ptr(v *uint32) null.Int {
	if v == nil {
		return null.NewInt(0, false)
	}
	if *v > math.MaxInt32 {
		panic("parameter value too large")
	}
	return null.IntFrom(int64(*v))
}

// TimeFromUnixInt64Ptr converts a Unix timestamp in seconds to null.Time.
func TimeFromUnixInt64Ptr(v *int64) null.Time {
	if v == nil {
		return null.NewTime(time.Time{}, false)
	}
	return null.TimeFrom(time.Unix(*v, 0))
}

// TimeFromUnixMilliInt64Ptr converts a Unix timestamp in milliseconds to null.Time.
func TimeFromUnixMilliInt64Ptr(v *int64) null.Time {
	if v == nil {
		return null.NewTime(time.Time{}, false)
	}
	return null.TimeFrom(time.UnixMilli(*v))
}

// TimeFromUnixMicroInt64Ptr converts a Unix timestamp in microseconds to null.Time.
func TimeFromUnixMicroInt64Ptr(v *int64) null.Time {
	if v == nil {
		return null.NewTime(time.Time{}, false)
	}
	return null.TimeFrom(time.UnixMicro(*v))
}
