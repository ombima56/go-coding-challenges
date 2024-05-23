package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	args1 := os.Args[1]
	args2 := os.Args[2]

	pos1 := 0
	pos2 := 0

	for pos2 < len(args2) {
		if pos1 < len(args1) && args1[pos1] == args2[pos2] {
			pos1++
		}
		pos2++

		if pos2 == len(args1) {
			for _, ch := range args1 {
				z01.PrintRune(ch)
			}
			z01.PrintRune('\n')
		}
	}
}
