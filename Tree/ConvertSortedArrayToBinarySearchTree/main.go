package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	l, r := 0, len(nums)-1
	m := l + (r-l)/2
	t := TreeNode{}

	t.Val = nums[m]
	t.Left = sortedArrayToBST(nums[l:m])
	t.Right = sortedArrayToBST(nums[m+1 : r+1])

	return &t
}
