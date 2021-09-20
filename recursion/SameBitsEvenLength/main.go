//https://www.geeksforgeeks.org/find-all-even-length-binary-sequences-with-same-sum-of-first-and-second-half-bits/
package main

import "fmt"

func sameEvenBits(n int) []string {
	res := []string{}

	var recur func(arr []rune)

	possible := []rune{'0', '1'}
	recur = func(arr []rune) {
		if len(arr) == n {
			copied := make([]rune, n*2)

			for i, c := range arr {
				copied[i] = c
				copied[i+n] = c
			}

			res = append(res, string(copied))
			return
		}

		for _, c := range possible {
			arr = append(arr, c)
			recur(arr)
			arr = arr[:len(arr)-1]
		}
	}

	recur([]rune{})

	return res
}

func main() {
	output := sameEvenBits(2)

	fmt.Println(output)
}
