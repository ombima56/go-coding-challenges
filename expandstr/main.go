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
	var word []rune
	var result []rune
	inWord := false

	for _, ch := range args {
		if ch != ' ' {
			word = append(word, ch)
			inWord = true
		} else if inWord {
			if len(result) > 0 {
				result = append(result, ' ', ' ', ' ')
			}
			result = append(result, word...)
			word = nil
			inWord = false
		}
	}

	if inWord {
		if len(result) > 0 {
			result = append(result, ' ', ' ', ' ')
		}
		result = append(result, word...)
	}

	for _, ch := range result {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
