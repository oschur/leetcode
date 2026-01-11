package main

func containsDuplicate(nums []int) bool {
	nHash := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if nHash[nums[i]] {
			return true
		}
		nHash[nums[i]] = true
	}
	return false
}
