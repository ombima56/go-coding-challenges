package main

// This function remove all spaces without importing strings.
func RemoveSpace(word string) string {
	var result string
	for _, ch := range word {
		if ch != ' ' {
			result += string(ch)
		}
	}
	return result
}
