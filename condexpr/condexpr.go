package condexpr

// Str 若expr成立，则返回a；否则返回b。
func Str(expr bool, a, b string) string {
	if expr {
		return a
	}
	return b
}

// Int 若expr成立，则返回a；否则返回b。
func Int(expr bool, a, b int) int {
	if expr {
		return a
	}
	return b
}

// Int8 若expr成立，则返回a；否则返回b。
func Int8(expr bool, a, b int8) int8 {
	if expr {
		return a
	}
	return b
}

// Int16 若expr成立，则返回a；否则返回b。
func Int16(expr bool, a, b int16) int16 {
	if expr {
		return a
	}
	return b
}

// Int32 若expr成立，则返回a；否则返回b。
func Int32(expr bool, a, b int32) int32 {
	if expr {
		return a
	}
	return b
}

// Int64 若expr成立，则返回a；否则返回b。
func Int64(expr bool, a, b int64) int64 {
	if expr {
		return a
	}
	return b
}

// Uint 若expr成立，则返回a；否则返回b。
func Uint(expr bool, a, b uint) uint {
	if expr {
		return a
	}
	return b
}

// Uint8 若expr成立，则返回a；否则返回b。
func Uint8(expr bool, a, b uint8) uint8 {
	if expr {
		return a
	}
	return b
}

// Uint16 若expr成立，则返回a；否则返回b。
func Uint16(expr bool, a, b uint16) uint16 {
	if expr {
		return a
	}
	return b
}

// Uint32 若expr成立，则返回a；否则返回b。
func Uint32(expr bool, a, b uint32) uint32 {
	if expr {
		return a
	}
	return b
}

// Uint64 若expr成立，则返回a；否则返回b。
func Uint64(expr bool, a, b uint64) uint64 {
	if expr {
		return a
	}
	return b
}

// Float32 若expr成立，则返回a；否则返回b。
func Float32(expr bool, a, b float32) float32 {
	if expr {
		return a
	}
	return b
}

// Float64 若expr成立，则返回a；否则返回b。
func Float64(expr bool, a, b float64) float64 {
	if expr {
		return a
	}
	return b
}
