package main

import (
	"fmt"
)

func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}

func RepeatAlpha(s string) string {
	var result string
	var num int
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			num = int(ch - 'a' + 1)
		} else if ch >= 'A' && ch <= 'Z' {
			num = int(ch - 'A' + 1)
		} else {
			result += string(ch)
		}

		for num > 0 {
			result += string(ch)
			num--
		}

	}
	return result
}
