// +build go1.18
//go:build go1.18

package condexpr

// Any 若expr成立，则返回a；否则返回b。
func Any[T any](expr bool, a, b T) T {
	if expr {
		return a
	}
	return b
}