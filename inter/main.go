package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	args1 := os.Args[1]
	args2 := os.Args[2]

	seen := make(map[rune]bool)
	exit := make(map[rune]bool)

	for _, ch1 := range args2 {
		exit[ch1] = true
	}
	
	for _, ch2 := range args1 {
		if exit[ch2] && !seen[ch2] {
			seen[ch2] = true
			z01.PrintRune(ch2)
		}
	}
	z01.PrintRune('\n')
}
