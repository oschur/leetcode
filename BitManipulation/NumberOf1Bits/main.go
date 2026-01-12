package main

func hammingWeight(n int) int {
	counter := 0
	for n != 0 {
		n = n & (n - 1)
		counter++
	}
	return counter
}
