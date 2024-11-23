package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
	"github.com/01-edu/z01"
)

func main() {
	table := [10]byte{}

	for j := 0; j < 5; j++ {
		for i := 0; i < random.IntBetween(7, 10); i++ {
			table[i] = byte(random.IntBetween(13, 126))
		}
		challenge.Function("PrintMemory", PrintMemory, solutions.PrintMemory, table)
	}
	table2 := [10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'}
	challenge.Function("PrintMemory", PrintMemory, solutions.PrintMemory, table2)
}

func PrintMemory(arr [10]byte) {
	for i, ch := range arr {
		printHex(ch)
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

func printHex(num byte) {
	hex := "0123456789abcdef"
	z01.PrintRune(rune(hex[num/16])) // Finding the most significant digit
	z01.PrintRune(rune(hex[num%16])) // finding the least significant digit
}
