package main

func maxProduct(nums []int) int {
	maxVal, minVal, res := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		k := nums[i]
		prevMin, prevMax := minVal, maxVal
		minVal = min(k, k*prevMin, k*prevMax)
		maxVal = max(k, k*prevMax, k*prevMax)
		res = max(res, maxVal)
	}
	return res
}
