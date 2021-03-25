package fib

func Fib(n int) int {
	fn := make(map[int]int)
	for i := 0; i <= n; i++ {
		f := 0
		if i <= 2 {
			f = 1
		} else {
			f = fn[i-1] + fn[i-2]
		}
		fn[i] = f

	}
	return fn[n]
}
