package main

import (
	"os"
)

func paramcount(s []string) int {
	count := len(s)
	return count
}

func main() {
	if len(os.Args) <= 1 {
		println(0)
		return
	}
	args := os.Args[1:]
	println(paramcount(args))
}
