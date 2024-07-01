package main

import (
	"fmt"
)

func LastWord(s string) string {
	check := false
	var words string
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			check = true
			words = string(s[i]) + words
		} else if check {
			break
		}
	}
	return words + "\n"
}

func main() {
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(LastWord(" lorem,ipsum "))
	fmt.Print(LastWord(" "))
}
