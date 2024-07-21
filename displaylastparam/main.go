package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) <= 1 {
		return
	}
	words := os.Args[1:]
	for _, word := range words[len(words)-1] {
		z01.PrintRune(word)
	}
	z01.PrintRune('\n')
}
