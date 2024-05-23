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
	check := false
	for _, c := range args {
		if c != ' ' {
			check = true
		}
	}
	if check {
		for i := len(args) - 1; i >= 0; i-- {

			if words == "" && string(args[i]) == " " {
				continue
			}
			if words != "" && string(args[i]) == " " {
				break
			}
			words = string(args[i]) + words

		}
		for _, ch := range words {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
}
