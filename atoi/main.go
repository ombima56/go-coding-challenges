package main

func Atoi(s string) int {
	var number int
	sign := 1

	for i, r := range s {
		if i == 0 && r == '+' {
			sign = 1
		} else if i == 0 && r == '-' {
			sign = 1
		} else if r >= '0' && r <= '9' {
			number = number*10 + int(r-'0')
		} else {
			return 0
		}
	}
	return sign * number
}
