package main

func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	prev2, prev1 := 1, 2
	for i := 3; i < n+1; i++ {
		prev2, prev1 = prev1, prev1+prev2
	}
	return prev1
}
