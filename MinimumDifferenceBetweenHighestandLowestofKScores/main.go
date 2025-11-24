package main

import "sort"

func main() {
	nums := []int{9, 4, 1, 7}
	k := 2
	minimumDifference(nums, k)
}

func minimumDifference(nums []int, k int) int {
	if k == 1 {
		return 0
	}

	sort.Ints(nums)
	diff := 100001
	for i := 0; i <= len(nums)-k; i++ {
		diff = min(diff, nums[i+k-1]-nums[i])
	}
	return diff
}
