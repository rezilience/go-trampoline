package trampoline

// Response is the struct used to facilitate the recursion, this is what your recursive function should return
type Response struct {
	Done   bool            // set it true when you hit the exit case of recursion
	Result interface{}     // this holds the return value of your recursive function
	Fn     func() Response // this should be set to your recursive function itself
}

// Trampoline avoids recursive stack by using a loop, thereby allowing large recursions without causing stack overflow
func Trampoline(fn func(args ...interface{}) Response) func(...interface{}) interface{} {
	return func(args ...interface{}) interface{} {
		r := fn(args...)
		for !r.Done {
			r = r.Fn()
		}
		return r.Result
	}
}
