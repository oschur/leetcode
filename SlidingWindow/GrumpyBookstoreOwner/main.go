package main

import "fmt"

func main() {
	customers := []int{1, 0, 1, 2, 1, 1, 7, 5}
	grumpy := []int{0, 1, 0, 1, 0, 1, 0, 1}
	minutes := 3
	fmt.Println(maxSatisfied(customers, grumpy, minutes))
}

func maxSatisfied(customers []int, grumpy []int, minutes int) int {

	allwaysHappy := 0
	extra := 0

	for i := 0; i < minutes; i++ {
		if grumpy[i] == 1 {
			extra += customers[i]
		} else {
			allwaysHappy += customers[i]
		}
	}
	maxExtra := extra

	for i := minutes; i < len(customers); i++ {
		if grumpy[i] == 1 {
			extra += customers[i]
		} else {
			allwaysHappy += customers[i]
		}
		if grumpy[i-minutes] == 1 {
			extra -= customers[i-minutes]
		}
		maxExtra = max(extra, maxExtra)
	}

	return allwaysHappy + maxExtra
}
