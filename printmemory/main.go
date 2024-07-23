package main

import (
	"unicode"

	"github.com/01-edu/z01"
)

// PrintMemory prints the memory of the byte array in hex format and ASCII representation.
func PrintMemory(arr [10]byte) {
	// Helper function to print a byte in hex format
	printHex := func(b byte) {
		hexDigits := "0123456789abcdef"
		z01.PrintRune(rune(hexDigits[b>>4]))
		z01.PrintRune(rune(hexDigits[b&0x0F]))
	}

	// Print the memory in hexadecimal format
	for i, b := range arr {
		printHex(b)
		if (i+1)%4 == 0 || i == len(arr)-1 {
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(' ')
		}
	}

	// Print the corresponding ASCII characters
	for _, b := range arr {
		if unicode.IsGraphic(rune(b)) {
			z01.PrintRune(rune(b))
		} else {
			z01.PrintRune('.')
		}
	}

	z01.PrintRune('\n')
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
