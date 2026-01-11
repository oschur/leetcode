package main

func twoSum(nums []int, target int) []int {
	nHash := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if num, ok := nHash[target-nums[i]]; ok {
			return []int{num, i}
		}
		nHash[nums[i]] = i
	}

	return []int{}
}
