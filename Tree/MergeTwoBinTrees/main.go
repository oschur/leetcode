package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	var resRoot TreeNode
	if root1 == nil && root2 == nil {
		return nil
	}

	if root1 == nil {
		resRoot.Val = root2.Val
		resRoot.Left = mergeTrees(nil, root2.Left)
		resRoot.Right = mergeTrees(nil, root2.Right)
		return &resRoot
	}

	if root2 == nil {
		resRoot.Val = root1.Val
		resRoot.Left = mergeTrees(root1.Left, nil)
		resRoot.Right = mergeTrees(root1.Right, nil)
		return &resRoot
	}

	resRoot.Val = root1.Val + root2.Val
	resRoot.Left = mergeTrees(root1.Left, root2.Left)
	resRoot.Right = mergeTrees(root1.Right, root2.Right)
	return &resRoot
}
