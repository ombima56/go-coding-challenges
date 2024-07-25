package main

import (
	"fmt"
)

func Lcm(first, second int) int {
	if second == 0 {
		return 0
	}
	return (first * second) / gcd(first, second)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	fmt.Println(Lcm(2, 7))
	fmt.Println(Lcm(0, 4))
}
