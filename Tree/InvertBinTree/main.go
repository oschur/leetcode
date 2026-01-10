package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTreeDFS(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	res := TreeNode{
		Val: root.Val,
	}

	res.Left = invertTreeDFS(root.Right)
	res.Right = invertTreeDFS(root.Left)
	return &res
}

func invertTreeBFS(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		node.Left, node.Right = node.Right, node.Left

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return root
}
