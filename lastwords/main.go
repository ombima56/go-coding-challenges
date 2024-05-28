package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	args := os.Args[1]

	var words string

	// checking if words is whitespaces.
	for i := len(args) - 1; i >= 0; i-- {
		if args[i] != ' ' {
			words = string(args[i]) + words
		} else {
			// If we encounter a space and we already have a word, break
			if words != "" {
				break
			}
		}
	}
		for _, ch := range words {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
}

