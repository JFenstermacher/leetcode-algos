package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructTree(nums []int) *TreeNode {
	return &TreeNode{}
}

func inorderTraversal(root *TreeNode) []int {
	res := []int{}

	var recur func(node *TreeNode)

	recur = func(node *TreeNode) {
		if node == nil {
			return
		}

		recur(node.Left)
		res = append(res, node.Val)
		recur(node.Right)
	}

	recur(root)

	return res
}

func main() {
}
