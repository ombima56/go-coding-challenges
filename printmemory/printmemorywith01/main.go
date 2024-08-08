package main

import (
	"github.com/01-edu/z01"
)

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

func PrintMemory(arr [10]byte) {
	for i, v := range arr {
		printStr(Hex(int(v)))
		if (i+1)%4 == 0 {
			z01.PrintRune('\n')
		} else if i != len(arr)-1 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')

	for _, ch := range arr {
		if ch >= 32 && ch <= 126 {
			z01.PrintRune(rune(ch))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')
}

func Hex(num int) string {
	var res string
	hex := "0123456789abcdef"

	if num == 0 {
		return "00"
	}
	for num > 0 {
		digit := num % 16
		res = string(hex[digit]) + res
		num /= 16
	}

	if len(res) == 1 {
		res = "0" + res
	}
	return res
}

func printStr(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
}
