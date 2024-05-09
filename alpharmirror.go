package main

func Alphamirror(s string) string {
	var result string
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			result += string('z' - (ch - 'a'))
		} else if ch >= 'A' && ch <= 'Z' {
			result += string('Z' - (ch - 'A'))
		} else {
			result += string(ch)
		}
	}
	return result
}
