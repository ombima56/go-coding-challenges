package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		os.Stdout.WriteString("0\n")
		return
	}
	var sum int
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}

	for i := num; i >= 2; i-- {
		if isPrime(i) {
			sum += i
		}

	}
	for _, ch := range strconv.Itoa(sum) {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')

}

func isPrime(nb int) bool {
	if nb <= 1 {
		return false
	}

	for i := 2; i*i <= nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
