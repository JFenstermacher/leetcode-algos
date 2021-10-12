package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructTree(nums []int) *TreeNode {
	return &TreeNode{}
}

func maxDepth(root *TreeNode) int {
	var recur func(node *TreeNode, count int) int

	recur = func(node *TreeNode, count int) int {
		if node == nil {
			return count - 1
		}

		left, right := recur(node.Left, count+1), recur(node.Right, count+1)

		if left > right {
			return left
		} else {
			return right
		}
	}

	return recur(root, 1)
}

func main() {
}
