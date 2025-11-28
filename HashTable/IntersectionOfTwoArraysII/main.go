package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println(intersection(nums1, nums2))
}

func intersection(nums1 []int, nums2 []int) []int {

	hashNum := make(map[int]int)
	resArr := make([]int, 0, min(len(nums1), len(nums2)))

	for i := 0; i < len(nums1); i++ {
		hashNum[nums1[i]]++
	}

	for i := 0; i < len(nums2); i++ {
		if hashNum[nums2[i]] > 0 {
			hashNum[nums2[i]]--
			resArr = append(resArr, nums2[i])
		}
	}

	return resArr
}
