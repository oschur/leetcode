package main

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	lower, upper := 0, n
	for lower <= upper {
		try := (lower + upper) / 2
		if isBadVersion(try) {
			upper = try - 1
		} else {
			lower = try + 1
		}
	}
	return upper + 1
}

func isBadVersion(x int) bool {
	return true
	// just for satisfying IDE
}
