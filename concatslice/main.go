package main

// import "fmt"

// func main() {
// 	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{4, 5, 6}))
// 	fmt.Println(ConcatSlice([]int{}, []int{4, 5, 6, 7, 8, 9}))
// 	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{}))
// }

import (
	// student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := [][][]int{
		{{1, 2, 3}, {4, 5, 6}},
		{{}, {-10, 0, 2}},
		{{-10, 0, 2}, {}},
		{{}, {}},
		{{1, 2, 3}, {4, 5, 6, 3, 4, 5, 6}},
		{{0, 0, 0}, {0, 0, 0}},
	}

	for _, arg := range args {
		challenge.Function("ConcatSlice", ConcatSlice, solutions.ConcatSlice, arg[0], arg[1])
	}
}

func ConcatSlice(slice1, slice2 []int) []int {
	if len(slice1) == 0 && len(slice2) == 0 {
		return []int(nil)
	}

	slice1 = append(slice1, slice2...)
	return slice1
}
