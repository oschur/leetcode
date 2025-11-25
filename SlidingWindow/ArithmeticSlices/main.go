package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 7, 9}
	fmt.Println(numberOfArithmeticSlices(nums))
}

func numberOfArithmeticSlices(nums []int) int {

	if len(nums) < 3 {
		return 0
	}

	counter, streak := 0, 0

	for i := 1; i < len(nums)-1; i++ {
		x, y, z := nums[i-1], nums[i], nums[i+1]
		if y-x == z-y {
			streak++
		} else {
			counter += triangleNum(streak)
			streak = 0
		}
	}
	counter += triangleNum(streak)

	return counter
}

func triangleNum(x int) int {
	return (x * (x + 1)) / 2
}
