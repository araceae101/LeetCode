package leetcode

// Problem: https://leetcode.com/problems/check-if-number-is-a-sum-of-powers-of-three/
// Author: Araceae
// Date: 2021/4/28
// Solution: Math
// Time Complexity: O(log3 N)
// Space Complexity: O(1) for iterative method and O(log3 N) for recursive method

// Problem: https://leetcode.com/problems/check-if-number-is-a-sum-of-powers-of-three/
func CheckPowersOfThreeIterative(n int) bool {
	r := 0
	for ; n > 0; n /= 3 {
		r = n % 3
		if r == 2 {
			return false
		}
	}
	return true
}

// Problem: https://leetcode.com/problems/check-if-number-is-a-sum-of-powers-of-three/
func CheckPowersOfThreeRecursive(n int) bool {
	return recursiveCheck(n, 0)
}

func recursiveCheck(n, r int) bool {
	if r == 2 {
		return false
	}
	if n == 0 {
		return true
	}
	return recursiveCheck(n/3, n%3)
}
