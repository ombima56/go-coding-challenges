// Write a program that displays its first argument, if there is one.
// Usage

// $ go run . hello there
// hello
// $ go run . "hello there" how are you
// hello there
// $ go run .
// $

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
	for _, word := range words[0]{
		z01.PrintRune(word)
	}
	z01.PrintRune('\n')
}
