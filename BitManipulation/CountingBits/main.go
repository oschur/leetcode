package main

func countBits(n int) []int {
	res := make([]int, n+1)
	offset := 1
	for i := 1; i < n+1; i++ {
		if offset*2 == i {
			offset = i
		}

		res[i] = res[i-offset] + 1
	}
	return res
}
