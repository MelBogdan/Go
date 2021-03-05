package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	x = fibb(x)
	fmt.Println(x)
}

func fibb(n int) int {
	m := map[int]int{}

	num := 1

	for i := 0; i <= n; i++ {
		if i == 0 || i == 1 {
			m[i] = num
			continue
		}
		m[i] = m[i-2] + m[i-1]
	}
	return m[n]
}
