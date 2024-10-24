package main

import (
	"fmt"
)

func main() {
	fmt.Println(FishAndChips(4))
	fmt.Println(FishAndChips(9))
	fmt.Println(FishAndChips(6))
}

func FishAndChips(n int) string {
	if n < 0 {
		return "error: number is negative"
	}

	switch {
	case n%2 == 0 && n%3 == 0:
		return "fish and chips"
	case n%2 == 0:
		return "fish"
	case n%3 == 0:
		return "chips"
	default:
		return "error: non divisible"
	}
}
