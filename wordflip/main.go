package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print(WordFlip("First second last"))
	fmt.Print(WordFlip(""))
	fmt.Print(WordFlip("     "))
	fmt.Print(WordFlip(" hello  all  of  you! "))
}

func WordFlip(str string) string {
	if str == "" {
		return "Invalid Output\n"
	}
	if len(str) == 0 {
		return "\n"
	}
	// Trim leading and trailing spaces
	str = strings.TrimSpace(str)

	words := strings.Fields(str)

	// Reverse the order of words
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ") + "\n"
}
