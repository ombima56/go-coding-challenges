package main

import (
	"strings"
)

// This functions reverse each words of a string, if the length of string is more than 4.
func SpinWords(str string) string {
	word := strings.Fields(str)
	var result string
	for _, words := range word {
		if len(words) > 4 {
			result += reverse(words) + " "
		} else {
			result += words + " "
		}
	}
	return strings.TrimSpace(result)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
