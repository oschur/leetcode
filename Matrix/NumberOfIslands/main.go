package main

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	counter := 0

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= m || c < 0 || c >= n || grid[r][c] != '1' {
			return
		}
		grid[r][c] = '0'

		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				counter++
				dfs(i, j)
			}
		}
	}
	return counter
}
