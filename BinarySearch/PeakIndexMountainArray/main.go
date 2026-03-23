package main

func peakIndexInMountainArray(arr []int) int {
	l, r := 0, len(arr)
	for r-l > 1 {
		m := l + (r-l)/2
		if good(arr, m) {
			l = m
		} else {
			r = m
		}
	}

	return l
}

func good(arr []int, i int) bool {
	if i == 0 || arr[i] > arr[i-1] {
		return true
	}
	return false
}
