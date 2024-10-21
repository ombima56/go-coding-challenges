package main

import (
	"os"

	"github.com/01-edu/z01"
)

// func main() {
// 	if len(os.Args) < 2 {
// 		return
// 	}

// 	args := os.Args[1:]
// 	for _, ch := range args {
// 		result := Brackets(ch)

// 		for _, ch := range result {
// 			z01.PrintRune(ch)
// 		}
// 		z01.PrintRune('\n')
// 	}
// }

// func Brackets(s string) string {
// 	stack := []rune{}
// 	pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}

// 	for _, c := range s {
// 		// Push opening brackets onto the stack
// 		if c == '(' || c == '[' || c == '{' {
// 			stack = append(stack, c)
// 		} else if c == ')' || c == ']' || c == '}' {
// 			// Check if there's a corresponding opening bracket in the stack
// 			if len(stack) == 0 || stack[len(stack)-1] != pairs[c] {
// 				return "Error"
// 			}
// 			// Pop the last opening bracket from the stack
// 			stack = stack[:len(stack)-1]
// 		}
// 	}

// 	// If stack is empty, all brackets matched correctly
// 	if len(stack) == 0 {
// 		return "OK"
// 	}
// 	return "Error"
// }

func main() {
	if len(os.Args) < 2 {
		return
	}

	for _, args := range os.Args[1:] {
		if checkBrackets(args) {
			printStr("OK")
		} else {
			printStr("Error")
		}
	}
}

func checkBrackets(s string) bool {
	stack := make([]rune, len(s))
	top := -1
	pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}

	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			top++
			stack[top] = c
		} else if c == ')' || c == ']' || c == '}' {
			if top < 0 || stack[top] != pairs[c] {
				return false
			}
			top--
		}
	}
	return top == -1
}

func printStr(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
