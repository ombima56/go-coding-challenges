package main

func BeZero(slice []int) []int {
	bezero := []int{}
	if slice == nil {
		return slice
	}
	seen := map[int]bool{}

	for _, ch := range slice {
		if !seen[ch] {
			seen[ch] = true
			bezero = append(bezero, ch)
		}
	}
	return bezero
}

// func BeZero(slice []int) []int {
// 	for i, _ := range slice {
// 		slice[i] = 0
// 	}
// 	return slice
// }
