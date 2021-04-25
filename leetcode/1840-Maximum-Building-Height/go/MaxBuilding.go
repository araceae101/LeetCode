package leetcode

// Problem: https://leetcode.com/problems/maximum-building-height/
// Author: Araceae
// Date: 2021/4/25
// Solution: Greedy
// Time Complexity: O(NlogN)
// Space Complexity: O(1) or O(N) depends on whether we can modify the original input array

import "sort"

// Problem: https://leetcode.com/problems/maximum-building-height/
func MaxBuilding(n int, A [][]int) int {
	A = append(A, []int{1, 0})
	A = append(A, []int{n, n - 1})

	sort.Slice(A, func(i, j int) bool { return A[i][0] < A[j][0] })

	m := len(A)

	for i := 1; i < m; i++ {
		A[i][1] = min(A[i][1], A[i-1][1]+A[i][0]-A[i-1][0])
	}
	for i := m - 2; i >= 0; i-- {
		A[i][1] = min(A[i][1], A[i+1][1]+A[i+1][0]-A[i][0])
	}

	res := 0
	for i := 1; i < m; i++ {
		h := max(A[i][1], A[i-1][1]) + (A[i][0]-A[i-1][0]-abs(A[i][1]-A[i-1][1]))/2
		res = max(res, h)
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
