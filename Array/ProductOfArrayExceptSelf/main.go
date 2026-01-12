package main

func productExceptSelf(nums []int) []int {
	n := len(nums)
	prefix, suffix := 1, 1
	resNums := make([]int, n)

	for i := 0; i < n; i++ {
		resNums[i] = prefix
		prefix *= nums[i]
	}

	for i := n - 1; i > 1; i-- {
		resNums[i] *= suffix
		suffix *= nums[i]
	}

	return resNums
}
