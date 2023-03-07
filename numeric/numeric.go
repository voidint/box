package numeric

type Number interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint | ~float32 | ~float64
}

// PositiveToPtr 若入参v大于零，则返回其指针；否则返回nil。
func PositiveToPtr[T Number](v T) *T {
	if v > 0 {
		return &v
	}
	return nil
}

// ExtractFromPtr 若入参指针为nil，则返回0；否则，返回指针指向的内存所存储的值。
func ExtractFromPtr[T Number](p *T) T {
	if p == nil {
		return 0
	}
	return *p
}
