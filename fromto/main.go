package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print(FromTo(1, 10))
	fmt.Print(FromTo(10, 1))
	fmt.Print(FromTo(10, 10))
	fmt.Print(FromTo(100, 10))
}

func FromTo(from int, to int) string {
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "Invalid\n"
	}

	result := ""

	// Determine the step direction based on the comparison of 'from' and 'to'
	if from <= to {
		for i := from; i <= to; i++ {
			// Convert number to string with leading zero if needed
			if i < 10 {
				result += "0" + strconv.Itoa(i)
			} else {
				result += strconv.Itoa(i)
			}
			if i != to {
				result += ", "
			}
		}
	} else {
		for i := from; i >= to; i-- {
			// Convert number to string with leading zero if needed
			if i < 10 {
				result += "0" + strconv.Itoa(i)
			} else {
				result += strconv.Itoa(i)
			}
			if i != to {
				result += ", "
			}
		}
	}

	// Append a newline character at the end of the result string
	result += "\n"
	return result
}
