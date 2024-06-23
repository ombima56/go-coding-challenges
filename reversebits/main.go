package main

import (
	"fmt"
)

func ReverseBits(oct byte) byte {
	var revbits byte = 0

	for i := 0; i < 8; i++ {
		revbits <<= 1
		revbits |= (oct & 1)
		oct >>= 1
	}
	return revbits
}

func main() {
	oct := byte(25)
	revbit := ReverseBits(oct)
	fmt.Printf("%08b||/", oct)
	fmt.Printf("%08b\n", revbit)
}
