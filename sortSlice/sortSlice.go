package sortSlice

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func SortSlice() {

	slice := createSlice()

	sortFunc(slice)

	fmt.Println("Your array", slice)
}

func createSlice() []int {

	var slice []int

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		line := scanner.Text()

		num, err := strconv.ParseInt(line, 10, 64)

		if err != nil {
			break
		}

		slice = append(slice, int(num))
	}

	return slice
}

func sortMy(slice []int) []int {
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

func sortFunc(slice []int) []int {
	sort.Ints(slice)
	return slice
}
