package main

// "github.com/01-edu/z01"

func LastRune(s string) rune {
	if len(s) <= 0 {
		return 0
	}
	runes := []rune(s)
	return runes[len(runes)-1]
}

// func main() {
// 	z01.PrintRune(LastRune("Hello!"))
// 	z01.PrintRune(LastRune("Salut!"))
// 	z01.PrintRune(LastRune("Ola!"))
// z01.PrintRune('\n')
// }
