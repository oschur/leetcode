package main

import "fmt"

func main() {
	s := "AAAAAAAAAAAAA"
	fmt.Println(findRepeatedDnaSequences(s))
}

func findRepeatedDnaSequences(s string) []string {
	res := make([]string, 0)

	if len(s) <= 10 {
		return res
	}

	hsh := make(map[string]int)

	for i := 0; i <= len(s)-10; i++ {
		hsh[string(s[i:i+10])]++
		if hsh[string(s[i:i+10])] == 2 {
			res = append(res, string(s[i:i+10]))
		}
	}
	return res
}
