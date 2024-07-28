package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("Invalid input.")
		return
	}
	arg1 := []rune(args[1])
	arg2 := []rune(args[2])

	if arg1[0] != ('[') || arg1[len(arg1)-1] != (']') {
		fmt.Println("Invalid input.")
		return
	}

	for _, n := range arg2 {
		if n < '0' || n > '9' {
			fmt.Println("Invalid target sum.")
			return
		}
	}

	for i, n := range arg1 {
		if i == 0 || i == len(arg1)-1 {
			continue
		}

		if n == ',' || n == ' ' {
			continue
		}

		if n == '-' {
			continue
		}

		if n < '0' || n > '9' {
			fmt.Printf("Invalid number: %s\n", string(n))
			return
		}

	}

	arr := []int{}
	sum, _ := strconv.Atoi(string(arg2))
	var str string

	for _, val1 := range arg1 {
		if val1 == '-' || (val1 >= '0' && val1 <= '9') {
			str += string(val1)
			fmt.Println(str)
		} else {
			num, _ := strconv.Atoi(str)
			str = ""
			if num != 0 {
				arr = append(arr, num)
			}
		}

	}

	result := [][]int{}
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == sum {
				result = append(result, []int{i, j})
			}
		}
	}

	if len(result) != 0 {
		fmt.Println("Pairs with sum", sum, ":", result)
	} else {
		fmt.Println("No pairs found.")
	}
}
