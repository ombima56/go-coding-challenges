package main

import (
	"os"
)

var Roman = []struct {
	value   int
	rNumber string
	working string
}{
	{1000, "M", "m"},
	{900, "CM", "(M-C)"},
	{500, "D", "D"},
	{400, "CD", "(D-C)"},
	{100, "C", "C"},
	{90, "XC", "(C-X)"},
	{50, "L", "L"},
	{40, "XL", "(L-X)"},
	{10, "x", "X"},
	{9, "IX", "(X-I)"},
	{5, "V", "V"},
	{4, "IV", "(V-I)"},
	{1, "I", "I"},
}

func main() {
	if len(os.Args) != 2 {
		return
	}

	roman, err := Atoi(os.Args[1])
	if roman <= 0 || roman >= 4000 {
		os.Stdout.WriteString("ERROR: cannot convert to roman digit\n")
		return
	}

	if err != "" {
		os.Stdout.WriteString(err)
		return
	}
	var working string
	var romanOutPut string

	for _, r := range Roman {
		for roman >= r.value {
			if working == "" {
				working += r.working
			} else {
				working += "+" + r.working
			}
			romanOutPut += r.rNumber
			roman -= r.value
		}
	}
	os.Stdout.WriteString(working + "\n" + romanOutPut + "\n")
}

func Atoi(s string) (int, string) {
	var number int
	sign := 1

	for i, v := range s {
		if i == 0 && v == '-' {
			sign = -1
		} else if i == 0 && v == '+' {
			sign = 1
		} else if v >= '0' && v <= '9' {
			number = number*10 + int(v-'0')
		} else {
			return 0, "ERROR: cannot convert to roman digit\n"
		}
	}
	return sign * number, ""
}
