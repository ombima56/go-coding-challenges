package main

import (
	"github.com/01-edu/z01"
)

func main() {
	result := []rune{}
	for i := 'z'; i >= 'a'; i-- {
		if (i-'z')%2 == 0 {
			result += string(i+ 32)
			z01.PrintRune(result)
		} else {
			z01.PrintRune(i)
		}
	}
	z01.PrintRune('\n')
}
