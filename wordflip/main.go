package main

import (
	"fmt"
)

func main() {
	fmt.Print(WordFlip("First second last"))
	fmt.Print(WordFlip(""))
	fmt.Print(WordFlip("     "))
	fmt.Print(WordFlip(" hello  all  of  you! "))
}

func WordFlip(str string) string {
	if len(str) == 0 {
		return "Invalid Input"
	}

	var words string
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] != ' ' {
			words += string(str[i])
		} else {
			words += string(str[i])
		}
	}
	return words + "\n"
}
