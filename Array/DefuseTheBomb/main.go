package main

import "fmt"

func main() {
	code := []int{5, 7, 1, 4}
	k := 3
	fmt.Println(decrypt(code, k))
}

func decrypt(code []int, k int) []int {

	n := len(code)
	codeRes := make([]int, n)

	for i := 0; i < n; i++ {
		sum := 0
		if k > 0 {
			for j := 1; j <= k; j++ {
				sum += code[(i+j)%n]
			}
			codeRes[i] = sum
		}
		if k < 0 {
			for j := 1; j <= k; j++ {
				sum += code[(i-j+n)%n]
			}
			codeRes[i] = sum
		}
	}
	return codeRes
}
