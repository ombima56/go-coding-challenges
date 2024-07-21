package main

import (
	"fmt"
	"strconv"
)

func ZipString(s string) string {
	var zip string
	var count = 1

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			zip += strconv.Itoa(count) + (string(s[i-1]))
			count = 1
		}
	}
	// Append the last counted charracters
	zip += strconv.Itoa(count) + string(s[len(s)-1])
	return zip
}

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}
