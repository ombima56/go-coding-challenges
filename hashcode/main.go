package main

import (
	"fmt"
)

func main() {
	table := []string{"Z", "Hi!", "BB198365", "sabito", "14 Avril 1912", "zyx987bca", "		pool-2020", "965truma747", " Mercedes-AMG GT"}
	for _, s := range table {
		fmt.Println(HashCode(s), solutions(s))
	}
}

func HashCode(dec string) (str string) {
	n := len(dec)

	for _, v := range dec {
		v = (v + rune(n)) % 127
		if v < 32 {
			v += 33
		}
		str += string(v)
	}
	return
}

func solutions(dec string) string {
	size := len(dec)
	hashed := ""
	for _, char := range dec {
		hash := (int(char) + size) % 127
		if hash < 32 || hash > 126 {
			hash += 33
		}
		hashed += string(hash)
	}
	return hashed
}
