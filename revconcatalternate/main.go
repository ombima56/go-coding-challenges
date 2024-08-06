package main

import (
	"log"
	"reflect"
)

var testCases = []struct {
	args [][]int
	want []int
}{
	{
		[][]int{{1, 2, 3}, {4, 5, 6}},
		[]int{3, 6, 2, 5, 1, 4},
	},
	{
		[][]int{{4, 5, 6, 7, 8, 9}, {1, 2, 3}},
		[]int{9, 8, 7, 6, 3, 5, 2, 4, 1},
	},
	{
		[][]int{{1, 2, 3}, {4, 5, 6, 7, 8, 9}},
		[]int{9, 8, 7, 3, 6, 2, 5, 1, 4},
	},
	{
		[][]int{{1, 2, 3}, {4, 5}},
		[]int{3, 2, 5, 1, 4},
	},
	{
		[][]int{{}, {4, 5, 6}},
		[]int{6, 5, 4},
	},
	{
		[][]int{{1, 2, 3}, {}},
		[]int{3, 2, 1},
	},
	{
		[][]int{{}, {}},

		[]int{},
	},
	{
		[][]int{{1, 2, 4}, {10, 20, 30, 40, 50}},

		[]int{50, 40, 4, 30, 2, 20, 1, 10},
	},
}

func main() {
	for _, tc := range testCases {
		got := RevConcatAlternate(tc.args[0], tc.args[1])
		if !reflect.DeepEqual(got, tc.want) {
			log.Fatalf("%s(%#v, %#v) == %#v instead of %#v\n",
				"RevConcatAlternate",
				tc.args[0],
				tc.args[1],
				got,
				tc.want,
			)
		}
	}
}

func RevConcatAlternate(slice1, slice2 []int) []int {
	// Reverse both slices
	reverse := func(slice []int) {
		for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
			slice[i], slice[j] = slice[j], slice[i]
		}
	}
	reverse(slice1)
	reverse(slice2)

	var result []int
	i, j := 0, 0

	// Interleave the elements
	for i < len(slice1) && j < len(slice2) {
		if i%2 == 0 {
			result = append(result, slice1[i])
		} else {
			result = append(result, slice2[j])
			j++
		}
		i++
	}

	// Append remaining elements from slice1
	for i < len(slice1) {
		result = append(result, slice1[i])
		i++
	}

	// Append remaining elements from slice2
	for j < len(slice2) {
		result = append(result, slice2[j])
		j++
	}

	return result
}

// func main() {
// 	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
// 	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
// 	fmt.Println(RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
// 	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{}))
// }
