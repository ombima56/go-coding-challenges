package main

import (
	"fmt"
)

func FifthAndSkip(str string) string {
	var count int
	var keep string
	if len(str) == 0 {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Input\n"
	}

	for _, ch := range str {
		if ch == ' ' {
			continue
		}
		
		if count == 5 {
			keep += " "
			count = 0
			continue
		}
		count++

		keep += string(ch)
	}
	return keep + "\n"
}

func main() {
	fmt.Print(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(FifthAndSkip("This is a short sentence"))
	fmt.Print(FifthAndSkip("1234"))
}
