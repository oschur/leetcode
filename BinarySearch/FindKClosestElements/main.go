package main

func findClosestElements(arr []int, k, x int) []int {
	l, r := -1, len(arr)-k
	for r-l > 1 {
		m := l + (r-l)/2
		if x-arr[m] > arr[m+k]-x {
			l = m
		} else {
			r = m
		}
	}
	return arr[r : r+k]
}
