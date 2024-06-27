package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	arg1 := os.Args[1]
	args2 := os.Args[2]
	value1, err1 := Atoi(arg1)
	value2, err2 := Atoi(args2)

	if err1 != nil || err2 != nil {
		return
	}
	answer := gcd(value1, value2)
	printStr(Itoa(answer))
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Atoi(s string) (int, error) {
	var num int
	sign := 1

	for i, ch := range s {
		if i == 0 && ch == '-' {
			sign = -1
		} else if i == 0 && ch == '+' {
			sign = 1
		} else if ch >= '0' && ch <= '9' {
			num = num*10 + int(ch-'0')
		} else {
			return 0, nil
		}
	}
	return sign * num, nil
}

func Itoa(n int) string {
	var num []rune
	for n > 0 {
		digit := n % 10
		num = append([]rune{rune('0' + digit)}, num...)
		n /= 10
	}
	return string(num)
}

func printStr(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
