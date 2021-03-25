package sortSlice

import (
	"testing"
)

var Slice []int

func BenchmarkSortMy(b *testing.B) {
	Slice = append(Slice, 4, 66, 7, 0, -102, 256, 1, 4)
	for i := 0; i < b.N; i++ {
		if Slice = sortMy(Slice); Slice[0] > Slice[1] {
			b.Fatalf("Uncorrect work of func")
		}
	}
}

func BenchmarkSortFunc(b *testing.B) {
	Slice = append(Slice, 4, 66, 7, 0, -102, 256, 1, 4)
	for i := 0; i < b.N; i++ {
		if Slice = sortFunc(Slice); Slice[0] > Slice[1] {
			b.Fatalf("Uncorrect work of func")
		}
	}
}

//BenchmarkSortMy-8        4033767               457 ns/op
//BenchmarkSortFunc-8       872598              1482 ns/op
//PASS
//ok      github.com/MelBogdan/Go/sortSlice       3.454s
