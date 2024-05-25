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

	runes := []rune(args)

	for i, ch := range runes {
		if ch >= 'a' && ch <= 'z' {
			runes[i] = 'a' + (ch-'a'+13)%26
		} else if ch >= 'A' && ch <= 'Z' {
			runes[i] = 'A' + (ch-'A'+13)%26
		}
	}
	for _, char := range runes {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
