package main

import "fmt"

func main() {
	fmt.Println(IsCapitalized("Hello! How are you?"))
	fmt.Println(IsCapitalized("Hello How Are You"))
	fmt.Println(IsCapitalized("Whats 4this 100K?"))
	fmt.Println(IsCapitalized("Whatsthis4"))
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(IsCapitalized(""))
}

func IsCapitalized(s string) bool {
	if s == "" || isSmall(rune(s[0])) {
		return false
	}

	for i := 1; i < len(s); i++ {
		if s[i] == ' ' && isSmall(rune(s[i+1])) {
			return false
		}
	}
	return true
}

func isSmall(s rune) bool {
	return s >= 'a' && s <= 'z'
}
