package main

import "fmt"

func main() {
	s := "xyzzaz"
	fmt.Println(countGoodSubstrings(s))
}

func countGoodSubstrings(s string) int {
	cnt := 0
	for i := 1; i < len(s)-1; i++ {
		if s[i-1] != s[i] && s[i] != s[i+1] && s[i-1] != s[i+1] {
			cnt++
		}
	}
	return cnt
}
