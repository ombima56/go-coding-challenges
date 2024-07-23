package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	hexdigits := "0123456789abcdef"
	result := ""

	args := os.Args[1]

	num, err := strconv.Atoi(args)
	if err != nil || num < 0 {
		for _, c := range "ERROR" {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
		return
	}

	if num == 0 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	for num > 0 {
		result = string(hexdigits[num%16]) + result
		num /= 16
	}

	for _, s := range result {
		z01.PrintRune(s)
	}
	z01.PrintRune('\n')
}
