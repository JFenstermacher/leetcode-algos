package main

import "fmt"

func Convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	rows := make([][]rune, numRows)

	index := 0
	dir := 1

	for _, c := range s {
		rows[index] = append(rows[index], c)

		index += dir

		if index == numRows {
			index -= 2
			dir = -1
		} else if index == -1 {
			index += 2
			dir = 1
		}
	}

	res := ""

	for _, row := range rows {
		res += string(row)
	}

	return res
}

func main() {
	input := "PAYPALISHIRING"
	output := Convert(input, 3)

	fmt.Println(output)
}
