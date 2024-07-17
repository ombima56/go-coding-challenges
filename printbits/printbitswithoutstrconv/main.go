package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	num, err := Atoi(os.Args[1])
	if err != nil {
		fmt.Print("00000000")
		return
	}

	var binaryStr string
	if num == 0 {
		binaryStr = "0"
	} else {
		for num > 0 {
			binaryStr = string(num%2+'0') + binaryStr
			num /= 2
		}
	}

	for len(binaryStr) < 8 {
		binaryStr = "0" + binaryStr
	}
	fmt.Print(binaryStr)
}

func Atoi(s string) (int, error) {
	var num int
	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			num = num*10 + int(ch-'0')
		} else {
			return 0, os.ErrInvalid
		}
	}
	return num, nil
}
