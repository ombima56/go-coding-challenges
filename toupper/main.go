package main

func ToUpper(s string) string {
	var upperCase string
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			char -= 32
		}
		upperCase += string(char)
	}
	return upperCase
}
