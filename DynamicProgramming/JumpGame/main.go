package main

func canJump(nums []int) bool {
	maxReach := 0
	for i := 0; i < len(nums); i++ {
		if i > maxReach {
			return false
		}
		if nums[i]+i > maxReach {
			maxReach = nums[i] + i
		}
	}
	return true
}
