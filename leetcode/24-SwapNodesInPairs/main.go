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

func swapPairs(head *ListNode) *ListNode {
	preHead := &ListNode{Next: head}
	curr := preHead

	for curr != nil {
		next, last := swap(curr.Next)

		curr.Next = next
		curr = last
	}

	return preHead.Next
}

func swap(node *ListNode) (*ListNode, *ListNode) {
	if node == nil {
		return nil, nil
	}

	first := node
	second := node.Next

	if second == nil {
		return first, nil
	}

	first.Next = second.Next
	second.Next = first

	return second, first
}

func recursiveSwapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	head.Next = recursiveSwapPairs(head.Next.Next)
	next.Next = head

	return next
}
