package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	var recur func(node *TreeNode, curr int) bool

	recur = func(node *TreeNode, curr int) bool {
		if node.Left == nil && node.Right == nil {
			return curr == targetSum
		}

		left, right := false, false

		if node.Left != nil {
			left = recur(node.Left, curr+node.Left.Val)
		}

		if node.Right != nil {
			right = recur(node.Right, curr+node.Right.Val)
		}

		return left || right
	}

	return recur(root, root.Val)
}

func main() {
}
