package main

import "fmt"

func romanToInt(s string) int {
	mappings := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	last := 0

	for _, c := range s {
		value := mappings[c]

		if value > last {
			total += (value - 2*last)
		} else {
			total += value
		}

		last = value
	}

	return total
}

func main() {
	output := romanToInt("MCMXCIV")

	fmt.Printf("%d", output)
}
