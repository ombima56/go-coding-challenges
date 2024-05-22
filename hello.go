// Write a program that displays "Hello World!" followed by a newline ('\n').
// Usage

// $ go run .
// Hello World!
// $

package main

import "github.com/01-edu/z01"

func main() {
	s := "Hello World!"
	for _, ch := range s {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
