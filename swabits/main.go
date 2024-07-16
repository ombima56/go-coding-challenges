package main

import (
	"fmt"
)

func SwapBits(octet byte) byte {
	return (octet << 4) | (octet >> 4)
}

func main() {
	// Example usage
	var input byte = 0b01000001 // 65 in decimal
	result := SwapBits(input)
	fmt.Printf("Swapped byte: %08b\n", result) // Output: Swapped byte: 00010000 (16 in decimal)
}
