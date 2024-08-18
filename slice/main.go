package main

import (
	"fmt"
)

func main() {
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
	fmt.Printf("%#v\n", Slice(a, 1))
	fmt.Printf("%#v\n", Slice(a, 2, 4))
	fmt.Printf("%#v\n", Slice(a, -3))
	fmt.Printf("%#v\n", Slice(a, -2, -1))
	fmt.Printf("%#v\n", Slice(a, 2, 0))
}

func Slice(a []string, nbrs ...int) []string {
	if len(nbrs) == 0 {
		return a
	}

	start := nbrs[0]
	if start < 0 {
		start += len(a)
	}

	if len(nbrs) == 1 {
		if start < 0 || start >= len(a) {
			return nil
		}
		return a[start:]
	}

	end := nbrs[1]
	if end < 0 {
		end += len(a)
	}

	// Special case: if end is 0, return nil
	if end == 0 {
		return nil
	}

	// Return empty slice if both indices are out of bounds but in correct order
	if start < 0 && end <= 0 && start < end {
		return []string{}
	}

	// Ensure start and end are within bounds
	if start < 0 {
		start = 0
	}
	if end > len(a) {
		end = len(a)
	}

	// Return nil if the range is invalid
	if start >= end || start >= len(a) {
		return nil
	}

	return a[start:end]
}
