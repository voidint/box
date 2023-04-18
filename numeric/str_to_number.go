package numeric

import "strconv"

// MustParseUint64 将字符串转换成uint。若转换过程中发生错误，则panic。
func MustParseUint64(s string) uint64 {
	v, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return v
}

// ParseUint64 将字符串转换成uint64。若转换过程中发生错误，则返回默认值。
func ParseUint64(s string, defVal uint64) uint64 {
	v, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return defVal
	}
	return v
}

// MustParseFloat64 将字符串转换成float64。若转换过程中发生错误，则panic。
func MustParseFloat64(s string) float64 {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return v
}

// ParseFloat64 将字符串转换成float64。若转换过程中发生错误，则丢弃错误并返回零值。
func ParseFloat64(s string, defVal float64) float64 {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return defVal
	}
	return v
}
