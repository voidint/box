package fn

import "runtime"

// FnName Returns the current function name
func FnName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}
