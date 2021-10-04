package main

import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < len(nums2); i++ {
		nums1[m+i] = nums2[i]
	}

	sort.Ints(nums1)
}

func main() {
	li := []int{1}
	li2 := []int{}

	merge(li, 1, li2, 0)

	fmt.Println(li)
}
