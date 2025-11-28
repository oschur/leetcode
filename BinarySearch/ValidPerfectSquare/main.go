package main

import "fmt"

func main() {
	n := 1
	fmt.Println(isPerfectSquare(n))
}

func isPerfectSquare(num int) bool {

	if num == 1 {
		return true
	}

	lower, upper := 0, num/2

	for lower <= upper {
		try := (lower + upper) / 2
		sq := try * try
		switch {
		case sq == num:
			return true
		case sq > num:
			upper = try - 1
		case sq < num:
			lower = try + 1
		}
	}
	return false
}
