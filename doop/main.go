package main

import (
	"os"
)

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}
	num := ""
	for n > 0 {
		digits := n % 10
		num = string(rune('0'+digits)) + num
		n /= 10
	}
	return sign + num
}

func Atoi(s string) (int, bool) {
	sign := 1
	num := 0

	for i, v := range s {
		if v == '-' && i == 0 {
			sign = -1
		} else if v == '+' && i == 0 {
			sign = 1
		} else if v >= '0' && v <= '9' {
			num = num*10 + int(v-'0')
		} else {
			return 0, false
		}
	}
	return sign * num, true
}

func PrintStr(s string) {
	os.Stdout.WriteString(s)
	os.Stdout.WriteString("\n")
}

func Calculation(a, operator, b string) string {
	value1, ok1 := Atoi(a)
	value2, ok2 := Atoi(b)

	if !ok1 || !ok2 {
		return ""
	}

	// Define max int64 value to handle large numbers
	const maxInt64 = 9223372036854775807

	if (value1 >= maxInt64 || value1 <= -maxInt64) || (value2 >= maxInt64 || value2 <= -maxInt64) {
		return ""
	}

	num := 0
	switch operator {
	case "+":
		num = value1 + value2
	case "-":
		num = value1 - value2
	case "*":
		num = value1 * value2
	case "/":
		if value2 == 0 {
			return "No division by 0"
		}
		num = value1 / value2
	case "%":
		if value2 == 0 {
			return "No modulo by 0"
		}
		num = value1 % value2
	default:
		return ""
	}
	return Itoa(num)
}

func main() {
	if len(os.Args) != 4 {
		return
	}
	value1, operator, value2 := os.Args[1], os.Args[2], os.Args[3]

	result := Calculation(value1, operator, value2)

	if result != "" {
		PrintStr(result)
	}
}