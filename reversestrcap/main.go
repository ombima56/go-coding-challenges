package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 2 {
		z01.PrintRune('\n')
		return
	}
	args := os.Args[1:]

	for i, arg := range args {
		printStr(formatWords(arg))
		if i < len(arg)-1 {
			printStr(" ")
		}
	}
	z01.PrintRune('\n')
}

func formatWords(s string) string {
	if len(s) == 0 {
		return s
	}

	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		if runes[i] >= 'A' && runes[i] <= 'Z' {
			runes[i] += 32
		}
	}

	if runes[len(runes)-1] >= 'a' && runes[len(runes)-1] <= 'z' {
		runes[len(runes)-1] -= 32
	}
	return string(runes)
}

func printStr(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
	// z01.PrintRune('\n')
}
