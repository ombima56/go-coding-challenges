package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error")
		return
	}

	input := os.Args[1]

	inputSlice := strings.Fields(input)

	var stack []int

	for _, numStr := range inputSlice {
		if num, err := strconv.Atoi(numStr); err == nil {
			stack = append(stack, num)
		} else {
			if len(stack) < 2 {
				fmt.Println("Error")
                return
			}

			// Pop the last two operands from the stack

			b, a := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var result int
			switch numStr {
				case "+":
                    result = a + b
                case "-":
                    result = a - b
                case "*":
                    result = a * b
                case "/":
                    if b == 0 {
                        fmt.Println("Error")
                        return
                    }
                    result = a / b
                default:
                    fmt.Println("Error")
                    return	
			}
			stack = append(stack, result)
		}
	}
	if len(stack) != 1 {
		fmt.Println("Error")
		return
	}
	fmt.Println(stack[0])
}
