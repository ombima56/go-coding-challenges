package main

import (
	"os"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 2 {
		z01.PrintRune('\n')
		return
	}
	args := os.Args[1:]

	for i, arg := range args {
		printStr(formatWords(arg))
		if i < len(arg)-1 {
			printStr(" ")
		}
	}
	z01.PrintRune('\n')

	{
		args := [][]string{
			{"First SMALL TesT"},
			{"SEconD Test IS a LItTLE EasIEr", "bEwaRe IT'S NoT HARd WhEN ", " Go a dernier 0123456789 for the road e"},
			{""},
		}

		for i := 0; i < 15; i++ {
			args = append(args, random.StrSlice(chars.Alnum))
		}

		for _, v := range args {
			challenge.Program("reversestrcap", v...)
		}
		challenge.Program("reversestrcap")
	}
}

func formatWords(s string) string {
	if len(s) == 0 {
		return s
	}

	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		if runes[i] >= 'A' && runes[i] <= 'Z' {
			runes[i] += 32
		}
	}

	if runes[len(runes)-1] >= 'a' && runes[len(runes)-1] <= 'z' {
		runes[len(runes)-1] -= 32
	}
	return string(runes)
}

func printStr(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
	// z01.PrintRune('\n')
}
