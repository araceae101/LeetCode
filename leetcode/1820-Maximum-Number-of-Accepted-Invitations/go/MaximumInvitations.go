package leetcode

func MaximumInvitations(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	match := make([]int, n)

	for v := 0; v < n; v++ {
		match[v] = -1
	}

	res := 0

	for u := 0; u < m; u++ {
		visited := make([]bool, n)
		for i := 0; i < n; i++ {
			visited[i] = false
		}
		if bmp(grid, visited, match, u) {
			res++
		}
	}

	return res
}

func bmp(grid [][]int, visited []bool, match []int, u int) bool {
	n := len(grid[0])

	for v := 0; v < n; v++ {
		if visited[v] || grid[u][v] == 0 {
			continue
		}

		visited[v] = true

		if match[v] < 0 || bmp(grid, visited, match, match[v]) {
			match[v] = u
			return true
		}
	}
	return false
}
