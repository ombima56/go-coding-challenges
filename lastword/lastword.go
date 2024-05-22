package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}

	arg1 := args[0]
	var word string
	// Find the last word
	for i := len(arg1) - 1; i >= 0; i-- {
		if arg1[i] != ' ' {
			word = string(arg1[i]) + word
		} else {
			// If we encounter a space and we already have a word, break
			if word != "" {
				break
			}
		}
	}
	// Print the last word
	for _, ch := range word {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
