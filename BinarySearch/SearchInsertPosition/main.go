package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 6}
	target := 4
	fmt.Println(searchInsert(nums, target))
}

func searchInsert(nums []int, target int) int {
	lower, upper := 0, len(nums)-1
	for lower <= upper {
		try := (upper + lower) / 2
		if nums[try] < target {
			lower = try + 1
		} else if nums[try] > target {
			upper = try - 1
		} else {
			return try
		}

	}
	return lower
}
