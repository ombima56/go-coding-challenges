package main

import "strings"

func Palidrome(str string) bool {
	var result string

	for _, ch := range str {
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') {
			result += strings.ToLower(string(ch))
		}
	}
	n := len(result)

	for i := 0; i < n/2; i++ {
		if result[i] != result[n-i-1] {
			return false
		}
	}
	return true
}
