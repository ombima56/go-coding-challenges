package main

import (
	"github.com/01-edu/z01"
)

func FoldInt(f func(int, int) int, a []int, n int) {
	for _, v := range a {
		n = f(n, v)
	}
	for _, ch := range Itoa(n) {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}

func Add(a, b int) int {
	return a + b
}

func Mul(a, b int) int {
	return a * b
}

func Sub(a, b int) int {
	return a - b
}

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	var digits []rune
	for n > 0 {
		digit := n % 10
		digits = append([]rune{rune('0' + digit)}, digits...)
		n /= 10
	}
	return string(digits)
}

// Example of a test function to test the functions.

// func main() {
// 	table := []int{1, 2, 3}
// 	ac := 93
// 	FoldInt(Add, table, ac)
// 	FoldInt(Mul, table, ac)
// 	FoldInt(Sub, table, ac)
// 	fmt.Println()

// 	table = []int{0}
// 	FoldInt(Add, table, ac)
// 	FoldInt(Mul, table, ac)
// 	FoldInt(Sub, table, ac)
// }
