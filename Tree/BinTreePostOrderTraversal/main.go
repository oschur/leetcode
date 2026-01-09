package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	if root.Left != nil {
		res = append(res, postorderTraversal(root.Left)...)
	}

	if root.Right != nil {
		res = append(res, postorderTraversal(root.Right)...)
	}

	res = append(res, root.Val)

	return res
}
