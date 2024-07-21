package main

import (
	"fmt"
)

func CanJump(slice []uint) bool {
	n := len(slice)
	if len(slice) == 0 {
		return false
	}

	if len(slice) == 1 {
		return true
	}

	var reachable int
	for i := 0; i < len(slice); i++ {
		// If the current index is beyond the furthest reachable index
		if i > reachable {
			return false
		}
		if int(i)+(int(slice[i])) > reachable {
			reachable = int(i) + int(slice[i])
		}
		// If reachable index is beyond or at the last index
		if reachable >= n-1 {
			return true
		}
	}
	// If we finish the loop and haven't reached the last index, return false
	return false
}

func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(CanJump(input1))

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(CanJump(input2))

	input3 := []uint{0}
	fmt.Println(CanJump(input3))
}
