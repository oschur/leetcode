package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}

	fmt.Println(lengthOfLis(nums))
}

func lengthOfLis(nums []int) int {
	sub := []int{nums[0]}

	for i := 1; i < len(nums); i++ {
		if nums[i] > sub[len(sub)-1] {
			sub = append(sub, nums[i])
		} else {
			k := sort.Search(len(sub), func(m int) bool { return sub[m] >= nums[i] })
			sub[k] = nums[i]
		}
	}

	return len(sub)
}
