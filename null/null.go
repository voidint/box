package null

import (
	"math"
	"time"

	"gopkg.in/guregu/null.v4"
)

// MustIntFromUint64Ptr 若入参为nil，则返回无效的null.Int；
// 若入参为正整数指针，则返回有效的null.Int；
// 若入参数值大于math.MaxInt64，则发生panic。
func MustIntFromUint64Ptr(v *uint64) null.Int {
	if v == nil {
		return null.NewInt(0, false)
	}
	if *v > math.MaxInt64 {
		panic("parameter value too large")
	}
	return null.IntFrom(int64(*v))
}

// MustIntFromUint32Ptr 若入参为nil，则返回无效的null.Int；
// 若入参为正整数指针，则返回有效的null.Int；
// 若入参数值大于math.MaxInt64，则发生panic。
func MustIntFromUint32Ptr(v *uint32) null.Int {
	if v == nil {
		return null.NewInt(0, false)
	}
	if *v > math.MaxInt32 {
		panic("parameter value too large")
	}
	return null.IntFrom(int64(*v))
}

// TimeFromUnixInt64Ptr 若入参为nil，则返回无效的null.Time；若入参非nil，则返回指定秒（Unix时间戳）的时间。
func TimeFromUnixInt64Ptr(v *int64) null.Time {
	if v == nil {
		return null.NewTime(time.Time{}, false)
	}
	return null.TimeFrom(time.Unix(*v, 0))
}

// TimeFromUnixMilliInt64Ptr 若入参为nil，则返回无效的null.Time；若入参非nil，则返回指定毫秒（Unix时间戳）的时间。
func TimeFromUnixMilliInt64Ptr(v *int64) null.Time {
	if v == nil {
		return null.NewTime(time.Time{}, false)
	}
	return null.TimeFrom(time.UnixMilli(*v))
}

// TimeFromUnixMicroInt64Ptr 若入参为nil，则返回无效的null.Time；若入参非nil，则返回指定微秒（Unix时间戳）的时间。
func TimeFromUnixMicroInt64Ptr(v *int64) null.Time {
	if v == nil {
		return null.NewTime(time.Time{}, false)
	}
	return null.TimeFrom(time.UnixMicro(*v))
}
