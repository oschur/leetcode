package main

import (
	"fmt"
)

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}

	fmt.Println(lengthOfLis(nums))
}

func lengthOfLis(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sub := []int{nums[0]}

	for i := 1; i < len(nums); i++ {
		if nums[i] > sub[len(sub)-1] {
			sub = append(sub, nums[i])
		} else {
			l, r := 0, len(sub)-1
			for l < r {
				m := l + (r-l)/2
				if nums[i] > sub[m] {
					l = m + 1
				} else {
					r = m
				}
			}
			sub[l] = nums[i]
		}
	}

	return len(sub)
}
