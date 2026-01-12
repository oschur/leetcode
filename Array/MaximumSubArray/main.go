package main

func maxSubArray(nums []int) int {
	sumNow, sumMax := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		sumNow = max(nums[i], sumNow+nums[i])
		sumMax = max(sumNow, sumMax)
	}
	return sumMax
}
