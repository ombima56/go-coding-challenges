
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Invalid input.")
		return
	}

	inputArr := os.Args[1]
	targetStr := os.Args[2]

	arr, err := parseArray(inputArr)
	if err != nil {
		fmt.Println(err)
		return
	}

	targetSum, err := strconv.Atoi(targetStr)
	if err != nil {
		fmt.Println("Invalid target sum.")
		return
	}

	result := findPairs(arr, targetSum)
	fmt.Println(result)
}

func parseArray(input string) ([]int, error) {
	input = strings.TrimSpace(input)
	if !strings.HasPrefix(input, "[") || !strings.HasSuffix(input, "]") {
		return nil, fmt.Errorf("Invalid input.")
	}

	input = input[1 : len(input)-1]
	if input == "" {
		return []int{}, nil
	}

	strElems := strings.Split(input, ", ")
	arr := make([]int, len(strElems))

	for i, elem := range strElems {
		elem = strings.TrimSpace(elem)
		num, err := strconv.Atoi(elem)
		if err != nil {
			return nil, fmt.Errorf("Invalid number: %s", elem)
		}
		arr[i] = num
	}

	return arr, nil
}

func findPairs(arr []int, targetSum int) string {
	pairs := [][2]int{}

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == targetSum {
				pairs = append(pairs, [2]int{i, j})
			}
		}
	}

	if len(pairs) == 0 {
		return "No pairs found."
	}

	return fmt.Sprintf("Pairs with sum %d: %v", targetSum, pairs)
}
