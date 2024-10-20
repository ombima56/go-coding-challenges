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
	result := ExandStr(args)

	for _, ch := range result {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}

func ExandStr(s string) string {
	spaceNeeded := false // Tracks if we need to add a space before a new word
	wordStarted := false // Tracks if we're inside a word
	// var result string
	var result []rune

	for _, ch := range s {
		if ch != ' ' {
			if spaceNeeded {
				result = append(result, ' ', ' ', ' ')
				// result += "   "
			}
			// result += string(ch)
			result = append(result, ch)
			wordStarted = true
			spaceNeeded = false
		} else if wordStarted {
			spaceNeeded = true // Set flag to insert space before the next word
		}
	}
	return string(result)
}
