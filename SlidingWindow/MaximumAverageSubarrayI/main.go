package main

import "fmt"

func main() {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4
	fmt.Println(findMaxAverage(nums, k))
}

func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	res := sum
	right := k
	for right < len(nums) {
		sum = sum - nums[right-k] + nums[right]
		if sum > res {
			res = sum
		}
		right++
	}
	return float64(res) / float64(k)
}
