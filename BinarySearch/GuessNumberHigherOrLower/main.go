package main

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	l, r := 0, n
	for r-l > 1 {
		m := (r + l) / 2
		if guess(m) != 1 {
			r = m
		} else {
			l = m
		}
	}
	return r

}

// Added so that the IDE does not give an error
func guess(num int) int {
	return num
}
