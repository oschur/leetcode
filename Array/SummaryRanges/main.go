package main

import "strconv"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	res := []string{}
	a, b := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			b = nums[i]
		} else {
			if a == b {
				res = append(res, strconv.Itoa(a))
			} else {
				res = append(res, strconv.Itoa(a)+"->"+strconv.Itoa(b))
			}
			a, b = nums[i], nums[i]
		}
	}

	if a == b {
		res = append(res, strconv.Itoa(a))
	} else {
		res = append(res, strconv.Itoa(a)+"->"+strconv.Itoa(b))
	}
	return res
}
