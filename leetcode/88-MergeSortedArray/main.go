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

func mergeRealSolution(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	for m > 0 && n > 0 {
		if nums1[m-1] < nums2[n-1] {
			nums1[m+n-1] = nums2[n-1]
			n--
		} else {
			nums1[m+n-1] = nums1[m-1]
			m--
		}
	}

	for i := 0; i < n; i++ {
		nums1[i] = nums2[i]
	}
}

func main() {
	li := []int{1}
	li2 := []int{}

	merge(li, 1, li2, 0)

	fmt.Println(li)
}
