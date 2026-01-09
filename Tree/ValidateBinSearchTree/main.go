package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return check(root, nil, nil)
}

func check(root *TreeNode, nmin, nmax *int) bool {
	if root == nil {
		return true
	}

	if nmin != nil && root.Val <= *nmin {
		return false
	}

	if nmax != nil && root.Val >= *nmax {
		return false
	}

	return check(root.Left, nmin, &root.Val) && check(root.Right, &root.Val, nmax)
}
