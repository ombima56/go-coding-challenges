// With this function for tabmult, you can  apply BasicAtoi concept instead of Atoi will output the same result.

// func BasicAtoi(s string) int{
// 	var numer int
// 	for _, v := range s {
// 		number = number*10 + int(v - '0')
// 	}
// 	return number
// }


package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	args := os.Args[1]
	num := Atoi(args)
	for i := 1; i <= 9; i++ {
		mult := i * num
		PrintValue(Itoa(i) + " " + "x" + " " + Itoa(num) + " " + "=" + " " + Itoa(mult)) 
	}

}

func Atoi(s string) int {
	var number int
	sign := 1
	for i, ch := range s {
		if i == 0 && ch == '-' {
			sign = -1
		} else if i == 0 && ch == '+' {
			sign = 1
		} else if ch >= '0' && ch <= '9' {
			number = number*10 + int(ch -'0')
		} else {
			return 0
		}
	}
	return sign * number
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

func PrintValue(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
