package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	res := TreeNode{
		Val: root.Val,
	}

	res.Left = invertTree(root.Right)
	res.Right = invertTree(root.Left)
	return &res
}
