package main

func pivotIndex(nums []int) int {
	sumNum := 0
	for i := 0; i < len(nums); i++ {
		sumNum += nums[i]
	}

	prefSum := 0
	for i := 0; i < len(nums); i++ {
		if sumNum-nums[i]-prefSum == prefSum {
			return i
		}
		prefSum += nums[i]
	}

	return -1
}
