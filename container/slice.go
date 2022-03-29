package container

// InSlice 返回目标元素是否在切片中的布尔值
func InSlice[T comparable](v T, items []T) bool {
	for i := range items {
		if v == items[i] {
			return true
		}
	}
	return false
}
