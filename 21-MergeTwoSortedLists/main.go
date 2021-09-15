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

func convertToArr(li *ListNode) []int {
	arr := []int{}

	for li != nil {
		arr = append(arr, li.Val)

		li = li.Next
	}

	return arr
}

func compareLists(l1 *ListNode, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}

		fmt.Printf("%d %d\n", l1.Val, l2.Val)
		l1 = l1.Next
		l2 = l2.Next
	}

	return l1 == nil && l2 == nil
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	curr := head

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = &ListNode{Val: l1.Val}
			l1 = l1.Next
		} else {
			curr.Next = &ListNode{Val: l2.Val}
			l2 = l2.Next
		}

		curr = curr.Next
	}

	if l1 != nil {
		curr.Next = l1
	}

	if l2 != nil {
		curr.Next = l2
	}

	return head.Next
}

func main() {
	l1 := convertToList([]int{1, 2, 4})
	l2 := convertToList([]int{1, 3, 4})

	expected := convertToList([]int{1, 1, 2, 3, 4, 4})
	output := mergeTwoLists(l1, l2)

	if compareLists(output, expected) {
		fmt.Println("Works")
	}
}
