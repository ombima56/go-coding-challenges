package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	args := os.Args[1:]
	union := make(map[rune]bool)

	for _, ch := range args[0] + args[1] {
		if !union[ch] {
			union[ch] = true
			z01.PrintRune(ch)
		}
	}
	z01.PrintRune('\n')
}
