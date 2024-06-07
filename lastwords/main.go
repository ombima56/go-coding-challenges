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
	check := false

	// checking if words is whitespaces.
	for i := len(args) - 1; i >= 0; i-- {
		if args[i] != ' ' {
			check = true
			words = string(args[i]) + words
		} else if check{
			break
		}
	}
	if len(words) > 0 {
		for _, ch := range words {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
}


