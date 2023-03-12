package str

// ExtractFromPtr 若入参指针为nil，则返回空字符串；否则，返回指针指向的内存所存储的字符串。
func ExtractFromPtr(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}
