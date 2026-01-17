package main

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}

	var backtrack func(start, remain int, track []int)
	backtrack = func(start, remain int, track []int) {
		if remain == 0 {
			n := make([]int, len(track))
			copy(n, track)
			res = append(res, n)
			return
		}

		for i := start; i < len(candidates); i++ {
			if candidates[i] > remain {
				break
			}
			backtrack(i, remain-candidates[i], append(track, candidates[i]))
		}
	}

	backtrack(0, target, nil)

	return res
}
