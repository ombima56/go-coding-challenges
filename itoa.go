package main

import (
	"github.com/01-edu/z01"
)

func Itoa(num int) string {
	if num == 0 {
		z01.PrintRune('0')
	}
	if num < 0 {
		z01.PrintRune('-')
		num = -num
	}
	var digits []rune
	for num > 0 {
		digit := num % 10
		digits = append([]rune{rune('0' + digit)}, digits...)
		num /= 10
	}
	return string(digits)
}
