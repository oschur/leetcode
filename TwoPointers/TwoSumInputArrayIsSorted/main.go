package main

func twoSum(numbers []int, target int) []int {
	resArr := make([]int, 2)
	l, r := 0, len(numbers)-1
	for l < r {
		sum := numbers[l] + numbers[r]
		if sum == target {
			resArr[0], resArr[1] = l+1, r+1
			return resArr
		} else if sum > target {
			r--
		} else {
			l++
		}
	}
	return resArr
}
