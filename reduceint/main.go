package main

import (
	"github.com/01-edu/z01"
)

func ReduceInt(a []int, f func(int, int) int) {
	if len(a) == 0 {
		return
	}
	reduce := a[0]
	for _, num := range a[1:] {
		reduce = f(num, reduce)
	}
	for _, ch := range Itoa(reduce) {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')

}

func Itoa(n int) string {
	if n == 0 {
		z01.PrintRune('0')

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

func main() {
	mul := func(acc int, cur int) int {
		return acc * cur
	}
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	div := func(acc int, cur int) int {
		return acc / cur
	}
	as := []int{500, 2}
	ReduceInt(as, mul)
	ReduceInt(as, sum)
	ReduceInt(as, div)
}

