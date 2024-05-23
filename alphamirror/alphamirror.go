package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	args := os.Args[1]

	var mirror string
	for _, ch := range args {
		if ch >= 'a' && ch <= 'z' {
			mirror += string('z' - (ch - 'a'))
		} else if ch >= 'A' && ch <= 'Z' {
			mirror += string('Z' - (ch - 'A'))
		} else {
			mirror += string(ch)
		}
	}
	for _, char := range mirror {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
