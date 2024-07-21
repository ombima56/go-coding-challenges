package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) <= 1 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}
	arg := os.Args[1:]
	count := '0'
	for range arg {
		count++
	}

	z01.PrintRune(count)
	z01.PrintRune('\n')

}

// func ParamCount()int  {
// var count int
// word := os.Args[1:]
// for range word {
// 	count++
// }
// return count
// }

// func PrintLn(num int) {
// 	for _, ch := range Itoa(num) {
// 		z01.PrintRune(ch)
// 	}
// 	z01.PrintRune('\n')
// }

// func Itoa(n int) string {
// 	if n == 0 {
// 		z01.PrintRune('0')
// 	}
// 	if n < 0 {
// 		z01.PrintRune('-')
// 		n = -n
// 	}
// 	var digits []rune
// 	for n > 0 {
// 		digit := n % 10
// 		digits = append([]rune{rune('0' + digit)}, digits...)
// 		n /= 10
// 	}
// 	return string(digits)
// }
