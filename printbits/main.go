package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	args := os.Args[1]

	num, err := strconv.Atoi(args)
	if err != nil {
		fmt.Print("00000000")
		return
	}

	var binaryStr string

	for num > 0 {
		if num%2 == 0 {
			binaryStr = "0" + binaryStr
		} else {
			binaryStr = "1" + binaryStr
		}
		num = num / 2
	}

	// Pad the binary string to 8 bits
	for len(binaryStr) < 8 {
		binaryStr = "0" + binaryStr
	}
	// If the number was 0, binaryStr will be empty
	if binaryStr == "" {
		binaryStr = "00000000"
	}
	fmt.Print(binaryStr)
}
