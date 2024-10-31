package main

import (
	"fmt"
)

func main() {
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatAlternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{}))
}

func ConcatAlternate(slice1, slice2 []int) []int {
	var result []int
	i, j := 0, 0
	len1, len2 := len(slice1), len(slice2)

	if len1 > len2 {
		for i < len1 && j < len2 {
			result = append(result, slice1[i], slice2[j])
			i++
			j++
		}
		// Append the remaining elements of slice1
		result = append(result, slice1[i:]...)
	} else if len2 > len1 {
		for i < len1 && j < len2 {
			result = append(result, slice2[j], slice1[i])
			i++
			j++
		}
		result = append(result, slice2[j:]...)
	} else {
		for i < len1 && j < len2 {
			result = append(result, slice1[i], slice2[j])
			i++
			j++
		}
	}
	return result
}
