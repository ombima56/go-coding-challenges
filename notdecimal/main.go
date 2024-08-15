package main

import "fmt"

func NotDecimal(dec string) string {
	n := len(dec)
	if n == 0 {
		return "\n"
	}

	var result string
	for _, ch := range dec {
		if ch >= 'a' && ch <= 'z' {
			return dec + "\n"
		}
		if ch == '.' {
			continue
		}
		result += string(ch)
	}
	for result[0] == '0' {
		result = result[1:]
	}
	return result + "\n"
}

func main() {
	fmt.Print(NotDecimal("0.1"))
	fmt.Print(NotDecimal("174.2"))
	fmt.Print(NotDecimal("0.1255"))
	fmt.Print(NotDecimal("1.20525856"))
	fmt.Print(NotDecimal("-0.0f00d00"))
	fmt.Print(NotDecimal(""))
	fmt.Print(NotDecimal("-19.525856"))
	fmt.Print(NotDecimal("1952"))
}
