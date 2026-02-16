package main

func subarraySum(nums []int, k int) int {
	pref := map[int]int{
		0: 1,
	}

	sumNums, count := 0, 0

	for i := 0; i < len(nums); i++ {
		sumNums += nums[i]
		count += pref[sumNums-k]
		pref[sumNums]++
	}

	return count
}
