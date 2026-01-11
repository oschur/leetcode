package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if subRoot == nil {
		return true
	}

	if root == nil {
		return false
	}

	if isIdentical(root, subRoot) {
		return true
	}

	return isSubtree(root.Right, subRoot) || isSubtree(root.Left, subRoot)
}

func isIdentical(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	if p.Val != q.Val {
		return false
	}

	return isIdentical(p.Left, q.Left) && isIdentical(p.Right, q.Right)
}
