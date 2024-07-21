package main

import (
	"strings"
)

// This function remove all spaces in a string.
func NoSpace(word string) string {
	return strings.ReplaceAll(word, " ", "")
}
