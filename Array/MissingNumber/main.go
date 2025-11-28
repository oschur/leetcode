package main

import "fmt"

func main() {
	nums := []int{3, 0, 1}
	fmt.Println(missingNumber(nums))
}

func missingNumber(nums []int) int {
	sumWithoutMissing := 0
	sumWithMissing := len(nums)
	for i := 0; i < len(nums); i++ {
		sumWithoutMissing += nums[i]
		sumWithMissing += i
	}
	return sumWithMissing - sumWithoutMissing
}
