package main

import "fmt"

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int)
	lenLon, left := 0, 0
	for right, c := range s {
		if prev, ok := m[c]; ok && prev >= left {
			if right-left > lenLon {
				lenLon = right - left
			}
			left = m[c] + 1
		}
		m[c] = right
	}
	if len(s)-left > lenLon {
		lenLon = len(s) - left
	}
	return lenLon
}
