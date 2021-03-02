package main

import "fmt"
import "math/big"

func main(){
	var n int64
	fmt.Println("Введите число n (1..n): ")
	fmt.Scan(&n)

	for i := int64(0); i < n; i++ {
		iBig  := big.NewInt(i)
		if iBig.ProbablyPrime(1) {
			fmt.Println(i)
		}
	}
}