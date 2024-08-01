package main

import "fmt"

func main() {
	PrintMemory([10]byte{'h'})
}

// PrintMemory prints the memory of the byte array in hex format and ASCII representation.
func PrintMemory(arr [10]byte) {
	for i, ch := range arr {
		if ch == '\x00' {
			fmt.Print("00")
		} else {
			fmt.Print(Hex(int(ch)))
		}
		if (i+1)%4 == 0 {
			fmt.Println()
		} else if i != len(arr)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()

	for _, ch := range arr {
		if ch >= 32 && ch <= 126 {
			fmt.Print(string(ch))
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}

func Hex(n int) string {
	hex := "0123456789abcdef"
	var str string
	if n == 0 {
		return "00"
	}

	for n > 0 {
		rem := n % 16
		str = string(hex[rem]) + str
		n /= 16
	}

	if len(str) == 1 {
		str = "0" + str
	}
	return str
}
