package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	adj := make([][]int, numCourses)
	for _, p := range prerequisites {
		course, pre := p[0], p[1]
		adj[pre] = append(adj[pre], course)
	}

	color := make([]int, numCourses)

	var dfs func(v int) bool
	dfs = func(v int) bool {
		if color[v] == 1 {
			return false
		}

		if color[v] == 2 {
			return true
		}

		color[v] = 1
		for _, neighbor := range adj[v] {
			if !dfs(neighbor) {
				return false
			}
		}

		color[v] = 2
		return true
	}

	for i := 0; i < numCourses; i++ {
		if color[i] == 0 {
			if !dfs(i) {
				return false
			}
		}
	}
	return true
}
