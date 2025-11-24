package main

import "fmt"

func main() {
	s := "cbaebabacd"
	p := "abc"
	fmt.Println(findAnagrams(s, p))
}

func findAnagrams(s string, p string) []int {

	if len(p) > len(s) {
		return nil
	}

	idxArr := make([]int, 0, len(s)/len(p))
	spHash := make(map[byte]int)

	for i := 0; i < len(p); i++ {
		spHash[p[i]]--
		spHash[s[i]]++
	}

	for i := 0; i < len(s)-len(p); i++ {

		if isAnagram(spHash) {
			idxArr = append(idxArr, i)
		}

		spHash[s[i]]--
		spHash[s[i+len(p)]]++
	}

	if isAnagram(spHash) {
		idxArr = append(idxArr, len(s)-len(p))
	}

	return idxArr
}

func isAnagram(spHash map[byte]int) bool {
	for _, chr := range spHash {
		if chr != 0 {
			return false
		}
	}
	return true
}
