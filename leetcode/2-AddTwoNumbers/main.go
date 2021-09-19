package main

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

func CarryOn(li *ListNode, carry int) *ListNode {
	if carry == 0 {
		return li
	}

	head := &ListNode{}
	curr := head

	for li != nil {
		value := li.Val + carry

		if value >= 10 {
			carry = 1
			value -= 10
		} else {
			carry = 0
		}

		curr.Next = &ListNode{Val: value}
		curr = curr.Next

		li = li.Next
	}

	if carry != 0 {
		curr.Next = &ListNode{Val: 1}
	}

	return head.Next
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: -1, Next: nil}
	curr := head
	carry := 0

	for l1 != nil && l2 != nil {
		value := l1.Val + l2.Val + carry

		if value >= 10 {
			carry = 1
			value -= 10
		} else {
			carry = 0
		}

		next := &ListNode{Val: value}
		curr.Next = next

		l1 = l1.Next
		l2 = l2.Next

		curr = curr.Next
	}

	if l1 == nil && l2 == nil {
		if carry == 1 {
			curr.Next = &ListNode{Val: 1}
		}
	} else if l1 == nil {
		curr.Next = CarryOn(l2, carry)
	} else if l2 == nil {
		curr.Next = CarryOn(l1, carry)
	}

	return head.Next
}
