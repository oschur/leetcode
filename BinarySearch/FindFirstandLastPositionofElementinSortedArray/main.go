package main

import "fmt"

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Println(searchRange(nums, target))
}

// bad solution

func badSearchRange(nums []int, target int) []int {
	resArr := []int{-1, -1}
	lower, upper := 0, len(nums)-1
	for lower <= upper {
		try := (lower + upper) / 2
		if nums[try] > target {
			upper = try - 1
		} else if nums[try] < target {
			lower = try + 1
		} else {
			swDown, swUp := try, try
			for swDown-1 >= 0 && nums[swDown] == nums[swDown-1] {
				swDown--
			}
			for swUp < len(nums)-1 && nums[swUp] == nums[swUp+1] {
				swUp++
			}
			resArr[0], resArr[1] = swDown, swUp
			break
		}
	}
	return resArr
}

// good solution

func searchRange(nums []int, target int) []int {
	return []int{findLeft(nums, target), findRight(nums, target)}
}

func findLeft(nums []int, target int) int {
	lower, upper := 0, len(nums)-1
	res := -1
	for lower <= upper {
		try := (lower + upper) / 2
		if nums[try] >= target {
			upper = try - 1
		} else {
			lower = try + 1
		}
		if nums[try] == target {
			res = try
		}
	}
	return res
}

func findRight(nums []int, target int) int {
	lower, upper := 0, len(nums)-1
	res := -1
	for lower <= upper {
		try := (lower + upper) / 2
		if nums[try] <= target {
			lower = try + 1
		} else {
			upper = try - 1
		}
		if nums[try] == target {
			res = try
		}
	}
	return res
}
