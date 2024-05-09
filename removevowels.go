package main

import (
	"strings"
)

func Disemvowel(comment string) string {
	found := map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'A': true, 'E': true, 'I': true, 'O': true, 'U': true}
	result := []string{}
	for _, ch := range comment {
		if !found[ch] {
			result = append(result, string(ch))
		}
	}
	return strings.Join(result, "")
}
