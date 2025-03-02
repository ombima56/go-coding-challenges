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
	var start, end = 0, len(a)

	if len(nbrs) > 0 {
		start = clampIndex(nbrs[0], len(a))
	}

	if len(nbrs) > 1 {
		end = clampIndex(nbrs[1], len(a))
	}

	if start > end {
		return nil
	}

	return a[start:end]
}

func clampIndex(idx, length int) int {
	if idx < 0 {
		idx += length
	}
	return max(0, min(idx, length))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
