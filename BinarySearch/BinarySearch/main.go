package main

import "fmt"

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 6
	fmt.Println(search(nums, target))
}

func search(nums []int, target int) int {
	lower, upper := 0, len(nums)-1
	for lower <= upper {
		mid := (upper + lower) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			upper = mid - 1
		} else {
			lower = mid + 1
		}
	}
	return -1
}
