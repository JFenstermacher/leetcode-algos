package main

import "fmt"

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}

	b = fmt.Sprintf("%0*s", len(a), b)

	res := make([]rune, len(a)+1)
	carry, back := 0, len(a)

	fmt.Println(a, b)
	for i := len(b) - 1; i >= 0; i-- {
		val := int(a[i]+b[i]) + carry - 96

		if val%2 == 0 {
			res[back] = '0'
		} else {
			res[back] = '1'
		}

		carry = val / 2
		back--
	}

	if carry == 1 {
		res[0] = '1'

		return string(res)
	} else {
		return string(res[1:])
	}
}

func main() {
	a := "111"
	b := "10"

	output := addBinary(a, b)

	fmt.Println(output)
}
