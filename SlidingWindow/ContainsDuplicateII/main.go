package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1, 2, 3}
	k := 2
	fmt.Println(containsNearbyDuplicate(nums, k))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	hash := make(map[int]int)
	for right := 0; right < len(nums); right++ {
		if i, ok := hash[nums[right]]; ok {
			if right-i <= k {
				return true
			}
		}
		hash[nums[right]] = right
	}
	return false
}
