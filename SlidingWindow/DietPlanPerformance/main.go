package main

import "fmt"

func main() {
	calories := []int{6, 5, 0, 0}
	k := 2
	lower := 1
	upper := 5
	fmt.Println(dietPlanPerformance(calories, k, lower, upper))
}

func dietPlanPerformance(calories []int, k int, lower int, upper int) int {
	sumCalories := 0
	pointCounter := 0
	for i := 0; i < k; i++ {
		sumCalories += calories[i]
	}

	updatePoints := func(sum int) {
		if sum < lower {
			pointCounter--
		} else if sum > upper {
			pointCounter++
		}
	}

	updatePoints(sumCalories)

	for right := k; right < len(calories); right++ {
		sumCalories = sumCalories - calories[right-k] + calories[right]
		updatePoints(sumCalories)
	}
	return pointCounter
}
