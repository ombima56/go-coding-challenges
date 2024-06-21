package main

import (
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	value1, err1 := Atoi(os.Args[1])
	operator := os.Args[2]
	value2, err2 := Atoi(os.Args[3])
	if !err1 || !err2 {
		return
	}

	var num int
	switch operator {
	case "+":
		num = value1 + value1
	case "-":
		num = value1 - value2
	case "/":
		if value2 == 0 {
			os.Stdout.WriteString("No division by 0\n")
			return
		}
		num = value1 / value2
	case "*":
		num = value1 * value2
	case "%":
		if value2 == 0 {
			os.Stdout.WriteString("No division by 0\n")
			return
		}
		num = value1 % value2
	default:
		return
	}

	if (value1 <= 9223372036854775807 && value1 >= 9223372036854775807) || (value2 <= 9223372036854775807 && value2 >= 9223372036854775807) {
		return
	}
	printStr(Itoa(num))
}

func Atoi(s string) (int, bool) {
	var number int
	sign := 1

	for i, ch := range s {
		if i == 0 && ch == '-' {
			sign = -1
		} else if i == 0 && ch == '+' {
			sign = 1
		} else if ch >= '0' && ch <= '9' {
			number = number*10 + int(ch-'0')
		} else {
			return 0, false
		}
	}
	return sign * number, true
}

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	if n < 0 {
		os.Stdout.WriteString("-")
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

func printStr(s string) {
	os.Stdout.WriteString(s + "\n")
}
