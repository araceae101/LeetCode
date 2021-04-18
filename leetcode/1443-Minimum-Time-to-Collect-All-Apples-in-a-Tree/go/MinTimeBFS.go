package leetcode

// Problem: https://leetcode.com/problems/minimum-time-to-collect-all-apples-in-a-tree/
// Author: Araceae
// Date: 2021/4/18

// Solution: BFS solution

// Time Complexity: O(N)
// Space Complexity: O(N)

// Performance:
// Runtime: 88 ms, faster than 100.00% of Go online submissions for Minimum Time to Collect All Apples in a Tree.
// Memory Usage: 17.4 MB, less than 100.00% of Go online submissions for Minimum Time to Collect All Apples in a Tree.

// Problem: https://leetcode.com/problems/minimum-time-to-collect-all-apples-in-a-tree/
func MinTimeBFS(n int, edges [][]int, hasApple []bool) int {

	nodeParent := make(map[int]int)

	for _, ed := range edges {
		_, ok := nodeParent[ed[1]]
		if !ok {
			nodeParent[ed[1]] = ed[0]
		} else {
			nodeParent[ed[0]] = ed[1]
		}
	}

	var q []int

	for i, a := range hasApple {
		if a {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		var nextq []int
		for _, currq := range q {
			parent := nodeParent[currq]

			// if the node has NOT yet visited, then add to the queue
			if !hasApple[parent] {
				hasApple[parent] = true
				nextq = append(nextq, parent)
			}
		}
		q = nextq
	}

	count := -1
	for _, a := range hasApple {
		if a {
			count++
		}
	}

	res := 0
	if count >= 0 {
		res = 2 * count
	}

	return res
}
