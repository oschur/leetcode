package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return compare(root.Left, root.Right)
}

func compare(x, y *TreeNode) bool {
	if x == nil || y == nil {
		return x == y
	}
	if x.Val != y.Val {
		return false
	}
	return compare(x.Left, y.Right) && compare(x.Right, y.Left)
}
