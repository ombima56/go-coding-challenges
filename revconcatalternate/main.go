package main

import "fmt"

func main() {
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{}))
}

func RevConcatAlternate(slice1, slice2 []int) []int {
	result := []int{}
	if len(slice2) == 0 {
		for i := len(slice1) - 1; i >= 0; i-- {
			result = append(result, slice1[i])
		}
		return result
	} else if len(slice2) > len(slice1) {
		for i := len(slice2) - 1; i >= len(slice1); i-- {
			result = append(result, slice2[i])
		}
		slice2 = slice2[:len(slice1)]
	} else if len(slice1) > len(slice2) {
		for i := len(slice1) - 1; i >= len(slice2); i-- {
			result = append(result, slice1[i])
		}
		slice1 = slice1[:len(slice2)]
	}

	for i := len(slice1) - 1; i >= 0; i-- {
		result = append(result, slice1[i], slice2[i])
	}

	return result
}
