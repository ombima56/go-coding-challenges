package main

import (
	"fmt"
)

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func FindPrevPrime(nb int) int {
	for i := nb - 1; i >= 2; i-- {
		if IsPrime(i) {
			return i
		}
	}
	return 0
}

func main() {
	fmt.Println(FindPrevPrime(28))
	fmt.Println(FindPrevPrime(6))
}
