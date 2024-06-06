package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main(){
	if len(os.Args) != 4 {
		return
	}
	args := os.Args[1:]
	var replWord string
	for _, ch := range args[0] {
		if string(ch) == args[1] {
			replWord += args[2]
		} else {
			replWord += string(ch)
		}
	}
	for _, ch := range replWord {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}