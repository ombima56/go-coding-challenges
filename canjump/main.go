package main

import (
	"fmt"
	"log"
	"reflect"
)

func main() {
	tests := []struct {
		args []uint
		want bool
	}{
		{args: []uint{2, 3, 1, 1, 4}, want: true},
		{args: []uint{1, 1, 1, 1, 0}, want: true},
		{args: []uint{5, 4, 3, 2, 1, 0}, want: true},
		{args: []uint{0}, want: true},
		{args: []uint{5}, want: true},
		{args: []uint{}, want: false},
		{args: []uint{1, 2, 3, 0, 2}, want: false},
		{args: []uint{3, 2, 1, 0, 4}, want: false},
		{args: []uint{0, 0, 0, 0, 0}, want: false},
		{args: []uint{1, 2, 3}, want: false},
		{args: []uint{1, 2, 3, 0, 1}, want: false},
		{args: []uint{1, 0, 0, 0, 0}, want: false},
		{args: []uint{1, 0, 1, 0, 1}, want: false},
		{args: []uint{10, 20, 30, 40, 0}, want: false},
	}

	for _, tc := range tests {
		got := CanJump(tc.args)
		if !reflect.DeepEqual(got, tc.want) {
			log.Fatalf("%s(%+v) == %+v instead of %+v\n",
				"CanJump",
				tc.args,
				got,
				tc.want,
			)
		}
	}
	fmt.Println("all test passed")
}

func CanJump(slice []uint) bool {

	if len(slice) == 0 {
		return false
	}

	if len(slice) == 1 {
		return true
	}

	var steps int
	for i := 0; i < len(slice); i++ {
		//if you encounter 0 in the slice you wont move so just return false
		if i > steps {
			return false
		}
		if steps == len(slice)-1 {
			return true
		} else {
			steps = i + int(slice[i])
		}
	}
	// If we finish the loop and haven't reached the last index, return false
	return false
}
