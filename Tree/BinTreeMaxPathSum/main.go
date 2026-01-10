package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxPath := -9999999999999
	deepSearch(root, &maxPath)
	return maxPath
}

func deepSearch(root *TreeNode, maxPath *int) int {
	if root == nil {
		return 0
	}

	pathSumLeft := max(deepSearch(root.Left, maxPath), 0)
	pathSumRight := max(deepSearch(root.Right, maxPath), 0)

	*maxPath = max(*maxPath, root.Val+pathSumLeft+pathSumRight)

	return root.Val + max(pathSumLeft, pathSumRight)
}
