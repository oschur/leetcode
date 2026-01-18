package main

func threeSum(nums []int) [][]int {
	res := [][]int{}
	quickSort(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		l, r := i+1, len(nums)-1
		sum := -nums[i]
		for l < r {
			if nums[l]+nums[r] == sum {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				r--
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			} else if nums[l]+nums[r] > sum {
				r--
			} else {
				l++
			}
		}
	}
	return res
}

func quickSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	left, right := 0, len(nums)-1

	for i := 0; i < right; i++ {
		if nums[i] < nums[right] {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		}
	}

	nums[left], nums[right] = nums[right], nums[left]
	quickSort(nums[:left])
	quickSort(nums[left+1:])
}
