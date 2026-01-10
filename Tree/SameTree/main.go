package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTreeDFS(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}
	return isSameTreeDFS(p.Left, q.Left) && isSameTreeDFS(p.Right, q.Right)
}

func isSameTreeBFS(p *TreeNode, q *TreeNode) bool {
	pQueue, qQueue := []*TreeNode{p}, []*TreeNode{q}

	for len(pQueue) > 0 && len(qQueue) > 0 {

		pNode, qNode := pQueue[0], qQueue[0]
		pQueue, qQueue = pQueue[1:], qQueue[1:]
		if pNode == nil && qNode == nil {
			continue
		}

		if pNode == nil || qNode == nil {
			return false
		}

		if pNode.Val != qNode.Val {
			return false
		}

		qQueue = append(qQueue, qNode.Left, qNode.Right)
		pQueue = append(pQueue, pNode.Left, pNode.Right)
	}

	return len(qQueue) == 0 && len(pQueue) == 0
}
