package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	expDepth := findTreeDepth(root)
	sum := addDeepestLeaveToSum(root, expDepth, 1)
	return sum
}

func findTreeDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(findTreeDepth(root.Left), findTreeDepth(root.Right))
}

func addDeepestLeaveToSum(root *TreeNode, expDepth, curDepth int) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil && curDepth == expDepth {
		return root.Val
	}

	return addDeepestLeaveToSum(root.Left, expDepth, curDepth+1) + addDeepestLeaveToSum(root.Right, expDepth, curDepth+1)
}
