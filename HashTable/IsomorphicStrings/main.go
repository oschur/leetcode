package main

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	cash1, cash2 := make(map[byte]int), make(map[byte]int)

	for i := 0; i < len(s); i++ {
		if cash1[s[i]] != cash2[t[i]] {
			return false
		}

		cash1[s[i]] = i + 1
		cash2[t[i]] = i + 1
	}
	return true
}
