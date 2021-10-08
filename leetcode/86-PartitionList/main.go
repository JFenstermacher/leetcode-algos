package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func convertToList(nums []int) *ListNode {
	head := &ListNode{}
	curr := head

	for _, el := range nums {
		curr.Next = &ListNode{Val: el}
		curr = curr.Next
	}

	return head.Next
}

func partition(head *ListNode, x int) *ListNode {
	above, below := &ListNode{}, &ListNode{}
	aboveHead, belowHead := above, below

	for head != nil {
		if head.Val < x {
			below.Next = head
			below = below.Next
		} else {
			above.Next = head
			above = above.Next
		}

		head = head.Next
	}

	above.Next = nil
	below.Next = aboveHead.Next

	return belowHead.Next
}

func main() {
	li := convertToList([]int{1, 4, 3, 2, 5, 2})

	output := partition(li, 3)

	for output != nil {
		fmt.Println(output.Val)

		output = output.Next
	}
}
