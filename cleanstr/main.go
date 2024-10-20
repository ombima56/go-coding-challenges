package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}

	args := os.Args[1]
	result := CleanStr(args)

	for _, ch := range result {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}

// func CleanStr(s string) string {
// 	var word []rune
// 	var result []rune
// 	inword := false
// 	for _, ch := range s {
// 		if ch != ' '  && ch != '\t' {
// 			word = append(word, ch)
// 			inword = true
// 		} else if inword {
// 			if len(result) > 0 {
// 				result = append(result, ' ')
// 			}
// 			result = append(result, word...)
// 			word = nil
// 			inword = false
// 		}
// 	}
// 	if inword {
// 		if len(result) > 0 {
// 			result = append(result, ' ')
// 		}
// 		result = append(result, word...)
// 	}
// 	return string(result)
// }

func CleanStr(s string) string {
	started := false     // Indicates if we've started appending characters
	spaceNeeded := false // Indicates if a space is needed between words
	result := ""

	for _, ch := range s {
		if ch != ' ' && ch != '\t' {
			if spaceNeeded {
				result += " " // Add space before a new word if needed
			}
			result += string(ch) // Append the current character to the result
			started = true
			spaceNeeded = false
		} else if started {
			// If we encounter a space after starting a word, mark space as needed
			spaceNeeded = true
		}
	}

	return result
}
