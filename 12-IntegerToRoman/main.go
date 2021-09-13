package main

import "strings"

func IntToRoman(num int) string {
	var b strings.Builder

	for num > 0 {
		if num >= 1000 {
			b.WriteString("M")
			num -= 1000
		} else if num >= 900 {
			b.WriteString("CM")
			num -= 900
		} else if num >= 500 {
			b.WriteString("D")
			num -= 500
		} else if num >= 400 {
			b.WriteString("CD")
			num -= 400
		} else if num >= 100 {
			b.WriteString("C")
			num -= 100
		} else if num >= 90 {
			b.WriteString("XC")
			num -= 90
		} else if num >= 50 {
			b.WriteString("L")
			num -= 50
		} else if num >= 40 {
			b.WriteString("XL")
			num -= 40
		} else if num >= 10 {
			b.WriteString("X")
			num -= 10
		} else if num >= 9 {
			b.WriteString("IX")
			num -= 9
		} else if num >= 5 {
			b.WriteString("V")
			num -= 5
		} else if num >= 4 {
			b.WriteString("IV")
			num -= 4
		} else {
			b.WriteString("I")
			num -= 1
		}
	}

	return b.String()
}
