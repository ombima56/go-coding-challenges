package main

import "fmt"

func HashCode(dec string) string {
	var result string
	size := len(dec)

	for _, char := range dec {
		hashedChar := (int(char) + size) % 127
		if hashedChar < 32 || hashedChar > 126 {
			hashedChar += 33
		}
		result += string(rune(hashedChar))
	}

	return result
}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}
