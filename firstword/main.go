package main

import (
	"fmt"
)

func FirstWord(s string) string {
	check := false
	var firstword string

	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			check = true
			firstword += string(s[i])
		} else if check {
			break
		}
	}
	return firstword + "\n"
}

func main() {
	fmt.Print(FirstWord("hello there"))
	fmt.Print(FirstWord(""))
	fmt.Print(FirstWord("hello   .........  bye"))
}
