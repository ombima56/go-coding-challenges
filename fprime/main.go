package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil || num <= 1 {
		return
	}

	factors := primeFactor(num)
	if len(factors) == 0 {
		return
	}
	fmt.Println(formFactors(factors))
}

func primeFactor(n int) []int {
	var factors []int
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}
	if n > 1 {
		factors = append(factors, n)
	}
	return factors
}

func formFactors(factors []int) string {
	var result string
	for i, factor := range factors {
		if i > 0 {

			result += "*"
		}
		result += strconv.Itoa(factor)
	}
	return result
}
