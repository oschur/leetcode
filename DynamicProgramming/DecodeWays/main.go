package main

func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}

	prev2, prev1 := 1, 1

	for i := 2; i <= len(s); i++ {
		temp := 0

		oneDigit := int(s[i-1] - '0')
		if oneDigit >= 1 && oneDigit <= 9 {
			temp += prev1
		}

		twoDigit := int(s[i-2]-'0')*10 + int(s[i-1]-'0')
		if twoDigit >= 10 && twoDigit <= 26 {
			temp += prev2
		}

		prev2 = prev1
		prev1 = temp
	}

	return prev1
}
