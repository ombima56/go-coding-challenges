package main

import "fmt"

func sumStr(a, b int) int {
	result := a + b
	return result
}

func main() {
	val1 := 3
	val2 := 5

	result := sumStr(val1, val2)
	fmt.Println(result)
}
