package leetcode

import "math"

// Problem: https://leetcode.com/problems/01-matrix/
// Author: Araceae
// Date: 2021/4/17

// Solution: DP solution
// to ask the neighbors in front for their nearest zero distance and plus 1:
//     (1) front (0, 0)     --to--> back  (m-1, n-1)
//     (2) back  (m-1, n-1) --to--> front (0, 0)

// Time Complexity: O(MN)
// Space Complexity: O(1)

// Runtime: 48 ms, faster than 98.32% of Go online submissions for 01 Matrix.
// Memory Usage: 6.5 MB, less than 96.64% of Go online submissions for 01 Matrix.

// Problem: https://leetcode.com/problems/01-matrix/
func UpdateMatrix(matrix [][]int) [][]int {

	m := len(matrix)
	if m < 1 {
		return [][]int{{}}
	}
	n := len(matrix[0])
	if n < 1 {
		return [][]int{{}}
	}

	// in-place solution
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				matrix[i][j] = math.MaxInt16
			}
		}
	}

	// chech from front (0, 0)  --to-->  back (m-1, n-1)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] != 0 {
				if i-1 >= 0 && matrix[i-1][j] != math.MaxInt16 && matrix[i][j] > matrix[i-1][j] {
					matrix[i][j] = matrix[i-1][j] + 1
				}
				if j-1 >= 0 && matrix[i][j-1] != math.MaxInt16 && matrix[i][j] > matrix[i][j-1] {
					matrix[i][j] = matrix[i][j-1] + 1
				}
			}
		}
	}

	// check from back (m-1, n-1)  --to-->  front (0, 0)
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if matrix[i][j] != 0 {
				if i+1 < m && matrix[i+1][j] != math.MaxInt16 && matrix[i][j] > matrix[i+1][j] {
					matrix[i][j] = matrix[i+1][j] + 1
				}
				if j+1 < n && matrix[i][j+1] != math.MaxInt16 && matrix[i][j] > matrix[i][j+1] {
					matrix[i][j] = matrix[i][j+1] + 1
				}
			}
		}
	}

	return matrix
}
