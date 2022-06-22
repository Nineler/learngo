package fib

//生成斐波那契数列(前两项为1，之后每项为前两项之和)
func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
