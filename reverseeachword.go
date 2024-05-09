package main

import (
	"strings"
)

func ReverseWords(str string) string {
	sliceArr := strings.Split(str, " ")
	s := ""
	v := []string{}
	for _, ch := range sliceArr {
		for _, char := range ch {
			s = string(char) + s
		}
		v = append(v, s)
		s = ""

	}
	return strings.Join(v, " ")
}
