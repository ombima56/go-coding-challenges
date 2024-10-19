package main

import "fmt"

func main() {
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
	fmt.Printf("%#v\n", Slice(a, 1))
	fmt.Printf("%#v\n", Slice(a, 2, 4))
	fmt.Printf("%#v\n", Slice(a, -3))
	fmt.Printf("%#v\n", Slice(a, -2, -1))
	fmt.Printf("%#v\n", Slice(a, 2, 0))
}

func Slice(a []string, nbrs ...int) []string {
	start := 0
	end := len(a)

	if len(nbrs) > 0 {
		start = nbrs[0]
		if start > len(a) {
			start = len(a)
		}
		if start < 0 {
			start = start + len(a)
		}
		if start < 0 {
			start = 0
		}
	}

	if len(nbrs) > 1 {
		end = nbrs[1]
		if end > len(a) {
			end = len(a)
		}
		if end < 0 {
			end = end + len(a)
		}
		if end < 0 {
			end = 0
		}
		if start > end {
			return nil
		}
	}

	return a[start:end]
}
