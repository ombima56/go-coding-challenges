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
	vowels := "aeiouAEIOU"
	if hasNoVowels(args, vowels) {
		printChar("No vowels")
		return
	}
	if isVowel(rune(args[0]), vowels) {
		printChar(args + "ay")
	} else {
		index := firstVowelIndex(args, vowels)
		if index != -1 {
			printChar(args[index:] + args[:index] + "ay")
		}
	}
}

func isVowel(ch rune, vowels string) bool {
	for _, v := range vowels {
		if ch == v {
			return true
		}
	}
	return false
}

func hasNoVowels(s, vowels string) bool {
	for _, ch := range s {
		if isVowel(ch, vowels) {
			return false
		}
	}
	return true
}

func firstVowelIndex(s, vowels string) int {
	for i, ch := range s {
		if isVowel(ch, vowels) {
			return i
		}
	}
	return -1
}

func printChar(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
