package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	args := os.Args[1]

	words := strings.Split(args, " ")

	var splitSpace []string
	for _, ch := range words {
		fmt.Println(string(ch))
		if ch != "" {
			splitSpace = append(splitSpace, ch)
		}
	}

	result := strings.Join(splitSpace, "   ")
	// fmt.Println(result)
	for _, ch := range result {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
