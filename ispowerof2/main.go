package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	args := os.Args[1]
	n := BAsicAtoi(args)

	if isPowerOf2(n) {
		printStr("true")
	} else {
		printStr("false")
	}
}

func BAsicAtoi(s string) int {
	var number int
	for _, ch := range s {
		number = number*10 + int(ch-'0')
	}
	return number
}

func isPowerOf2(n int) bool {
	if n <= 0 {
		return false
	}
	return (n & (n - 1)) == 0
}

func printStr(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
