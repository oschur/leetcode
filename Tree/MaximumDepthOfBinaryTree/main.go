package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepthByDfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepthByDfs(root.Left)
	right := maxDepthByDfs(root.Right)
	if left > right {
		return 1 + left
	}
	return 1 + right
}

func maxDepthByBfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 0

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		depth++

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return depth
}
