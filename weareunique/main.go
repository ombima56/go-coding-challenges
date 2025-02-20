package main

import (
	"fmt"
)

func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}

func WeAreUnique(str1 , str2 string) int {
	if len(str1) == 0 && len(str2) == 0 {
		return -1
	}

	// Maps to store unique characters in each string
	var chars1 = make(map[rune]bool)
	var chars2 = make(map[rune]bool)

	// Populate chars1 with characters from str1
	for _, ch := range str1 {
		chars1[ch] = true
	}

	// Populate chars2 with characters from str2
	for _, ch := range str2 {
		chars2[ch] = true
	}

	var uniqueCount int

	// Count characters in str1 that are NOT in str2
	for char := range chars1 {
		if !chars2[char] {
			uniqueCount++
		}
	}

	// Count characters in str2 that are NOT in str1
	for char := range chars2 {
		if !chars1[char] {
			uniqueCount++
		}
	}

	return uniqueCount
}
