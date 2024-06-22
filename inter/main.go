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

	// Maps to track seen characters and characters existing in the second string
	seen := make(map[rune]bool)
	exists := make(map[rune]bool)

	// Populate the exists map with characters from the second string
	for _, ch := range args2 {
		exists[ch] = true
	}

	for _, ch := range args1 {
		// If the character exists in the second string and has not been seen before
		if exists[ch] && !seen[ch] {
			seen[ch] = true
			z01.PrintRune(ch)
		}
	}
	z01.PrintRune('\n')
}
