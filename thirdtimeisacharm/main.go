package main

import (
	"fmt"
)

func main() {
	fmt.Print(ThirdTimeIsACharm("123456789"))
	fmt.Print(ThirdTimeIsACharm(""))
	fmt.Print(ThirdTimeIsACharm("a b c d e f"))
	fmt.Print(ThirdTimeIsACharm("12"))
}

func ThirdTimeIsACharm(str string) string {
	if len(str) == 3 {
		return "\n"
	}

	var charm string

	for i := 0; i < len(str); i++ {
		if (i+1)%3 == 0 {
			charm += string(str[i])
		}
	}
	return charm + "\n"
}
