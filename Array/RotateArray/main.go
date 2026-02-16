package main

func rotate(nums []int, k int) {
	n := len(nums)
	k %= n
	rotateHelper(nums, 0, n)
	rotateHelper(nums, 0, k)
	rotateHelper(nums, k, n)
}

func rotateHelper(nums []int, i, j int) {
	j--
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		j--
		i++
	}
}
