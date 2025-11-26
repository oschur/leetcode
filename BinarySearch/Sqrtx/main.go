package main

import "fmt"

func main() {
	x := 5
	fmt.Println(mySqrt(x))
}

func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	lower, upper := 1, x/2
	for lower <= upper {
		mid := (lower + upper) / 2
		square := mid * mid
		if square == x {
			return x
		}
		if square < x {
			lower = mid + 1
		}
		if square > x {
			upper = mid - 1
		}
	}
	return upper
}
