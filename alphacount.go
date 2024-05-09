package main

func AlphaCount(s string) int {
	var count int
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' {
			count++
		}
	}
	return count
}
