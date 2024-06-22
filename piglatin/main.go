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
	result := pigLatin(args)

	for _, ch := range result {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}

func isVowel(ch byte) bool {
	return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' || ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U'
}

func pigLatin(word string) string {

	//  Check if the word start with a vowel
	if isVowel(word[0]) {
		return word + "ay"
	}

	// Find the position of the first vowel
	for i := 0; i < len(word); i++ {
		if isVowel(word[i]) {
			return word[i:] + word[:i] + "ay"
		}
	}
	return "No Vowels"
}
