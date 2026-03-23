package main

func search(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1
	for l <= r {
		m := l + (r-l)/2
		if nums[m] == target {
			return m
		}
		if nums[m] < nums[r] {
			if target > nums[m] && target <= nums[r] {
				l = m + 1
			} else {
				r = m - 1
			}
		} else {
			if target < nums[m] && target >= nums[l] {
				r = m - 1
			} else {
				l = m + 1
			}
		}
	}
	return -1
}

func findOffset(nums []int) int {
	l, r := -1, len(nums)-1
	for r-l > 1 {
		m := l + (r-l)/2
		if nums[m] > nums[len(nums)-1] {
			l = m
		} else {
			r = m
		}
	}
	return r
}

func searchAnother(nums []int, target int) int {
	offset := findOffset(nums)
	l, r := 0, len(nums)

	for r-l > 1 {
		m := l + (r-l)/2
		mid := (m + offset) % len(nums)
		if nums[mid] <= target {
			l = m
		} else {
			r = m
		}
	}

	realLeft := (l + offset) % len(nums)
	if nums[realLeft] == target {
		return realLeft
	}
	return -1
}
