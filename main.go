package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	slice := []int{}

	createSlice(slice)

	slice = sort(slice)

	fmt.Println("Your array", slice)
}

func createSlice(slice []int) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		line := scanner.Text()

		num, err := strconv.ParseInt(line, 10, 64)

		if err != nil {
			break
		}

		slice = append(slice, int(num))
	}
}

func sort(slice []int) []int {
	var n = len(slice)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if slice[j-1] > slice[j] {
				slice[j-1], slice[j] = slice[j], slice[j-1]
			}
			j = j - 1
		}
	}
	return slice
}
