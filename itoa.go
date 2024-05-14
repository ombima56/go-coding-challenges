package main

func Itoa(num int) string {
	var number string
	sign := 1
	var mod int
	var div int
	var numSlice []int

	if num == 0 {
		return "0"
	}
	if num < 0 {
		sign = num * -1
		number = Itoa(sign)
		number = "-" + number
		return number
	}
	div = num
	for div > 0 {
		mod = div % 10
		div = div / 10
		numSlice = append(numSlice, mod)
	}
	for i := len(numSlice) - 1; i >= 0; i-- {
		number = number + string(int32(numSlice[i])+'0')
	}
	return number

}
