package main

import "fmt"

func main() {
	nums := []int{1, 1, 0, 1}
	fmt.Println(longestSubarray(nums))
}

func longestSubarray(nums []int) int {
	left := 0
	counterZeroes := 0
	lenLon := 0

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			counterZeroes++
		}

		for counterZeroes > 1 {
			if nums[left] == 0 {
				counterZeroes--
			}
			left++
		}
		lenLon = max(lenLon, right-left)
	}
	return lenLon
}
