package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 240
	k := 2
	fmt.Println(divisorSubstrings(num, k))
}

func divisorSubstrings(num int, k int) int {
	strnum := strconv.Itoa(num)
	counter := 0

	for i := 0; i < len(strnum)-k+1; i++ {
		m, _ := strconv.Atoi(strnum[i : i+k])

		if isDivisor(num, m) {
			counter++
		}
	}

	return counter
}

func isDivisor(n int, m int) bool {
	if m != 0 && n%m == 0 {
		return true
	}
	return false
}
