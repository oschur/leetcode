package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	prefix := make(map[int]int)
	prefix[0]++

	var dfs func(node *TreeNode, curSum int) int
	dfs = func(node *TreeNode, curSum int) int {
		if node == nil {
			return 0
		}

		curSum += node.Val

		countPaths := prefix[curSum-targetSum]

		prefix[curSum]++
		left := dfs(node.Left, curSum)
		right := dfs(node.Right, curSum)
		prefix[curSum]--

		return countPaths + left + right
	}
	return dfs(root, 0)
}
