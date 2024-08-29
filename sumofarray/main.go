package main

func FindEvenIndex(arr []int) int {

	for i := 0; i < len(arr); i++ {
		var indexLeft int
		var indexRight int
		for j := 0; j < i; j++ {
			indexLeft += arr[j]
		}

		for j := i + 1; j < len(arr); j++ {
			indexRight += arr[j]
		}

		if indexLeft == indexRight {
			return i
		}
	}
	return -1
}
