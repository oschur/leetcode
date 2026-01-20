package main

func pacificAtlantic(heights [][]int) [][]int {
	m := len(heights)
	n := len(heights[0])

	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	var dfs func(r, c int, visited [][]bool)
	dfs = func(r, c int, visited [][]bool) {
		visited[r][c] = true
		for _, d := range dirs {

			nr, nc := r+d[0], c+d[1]
			if nr < 0 || nr >= m || nc < 0 || nc >= n {
				continue
			}

			if visited[nr][nc] {
				continue
			}

			if heights[nr][nc] < heights[r][c] {
				continue
			}

			dfs(nr, nc, visited)
		}
	}

	pac, atl := make([][]bool, m), make([][]bool, m)
	for i := 0; i < m; i++ {
		pac[i] = make([]bool, n)
		atl[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		dfs(i, 0, pac)
	}
	for j := 0; j < n; j++ {
		dfs(0, j, pac)
	}

	for i := 0; i < m; i++ {
		dfs(i, n-1, atl)
	}
	for j := 0; j < n; j++ {
		dfs(m-1, j, atl)
	}

	var res [][]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if atl[i][j] && pac[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}
