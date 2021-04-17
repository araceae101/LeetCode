package leetcode

// Problem: https://leetcode.com/problems/01-matrix/
// Author: Araceae
// Date: 2021/4/17

// Solution: BFS solution
// Like Ripples, start from Zero-element and tell how they are close to Zero to their 4 neighbors
// and the 4 neighbors continue to tell their neighbors ... until all the elements all get the signal.

// Time Complexity: O(MN)
// Space Complexity: O(MN)

// Performance:
// Runtime: 52 ms, faster than 96.64% of Go online submissions for 01 Matrix.
// Memory Usage: 6.9 MB, less than 90.76% of Go online submissions for 01 Matrix.

// Problem: https://leetcode.com/problems/01-matrix/
func UpdateMatrixBFS(matrix [][]int) [][]int {
	if matrix == nil {
		return [][]int{{}}
	}
	m := len(matrix)
	if m < 1 {
		return [][]int{{}}
	}
	n := len(matrix[0])
	if n < 1 {
		return [][]int{{}}
	}

	var queue [][2]int
	dir := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				matrix[i][j] = -1
			} else {
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	for len(queue) > 0 {
		var nextQueue [][2]int
		for _, q := range queue {
			for _, d := range dir {
				x := q[0] + d[0]
				y := q[1] + d[1]

				if x < 0 || y < 0 || x >= m || y >= n {
					continue
				}

				if matrix[x][y] == -1 {
					matrix[x][y] = matrix[q[0]][q[1]] + 1
					nextQueue = append(nextQueue, [2]int{x, y})
				}
			}
		}
		queue = nextQueue
	}

	return matrix
}
