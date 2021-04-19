package leetdocde

// Problem: https://leetcode.com/problems/combination-sum-iv/
// Author: Araceae
// Date: 2021/4/19

// Solution: DP

// Time Complexity: O(N*target)
// Space Complexity: O(target)

// Performance:
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Combination Sum IV.
// Memory Usage: 2 MB, less than 100.00% of Go online submissions for Combination Sum IV.

// Problem: https://leetcode.com/problems/combination-sum-iv/
func CombinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1

	for t := 1; t <= target; t++ {
		for _, n := range nums {
			if t-n >= 0 {
				dp[t] += dp[t-n]
			}
		}
	}

	return dp[target]
}
