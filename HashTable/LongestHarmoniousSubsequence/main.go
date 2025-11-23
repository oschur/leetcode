package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 1, 3, -5}
	fmt.Println(findLHS(nums))
}

func findLHS(nums []int) int {
	hmap := countOccurances(nums)
	res := 0

	for num, cnt := range hmap {
		if hmap[num+1] > 0 {
			res = max(res, cnt+hmap[num+1])
		}
	}
	return res
}

func countOccurances(nums []int) map[int]int {
	hmap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		hmap[nums[i]]++
	}
	return hmap
}
