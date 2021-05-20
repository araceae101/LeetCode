package leetcode

import "github.com/araceae101/LeetCode/pkg/unittest"

/**
 * Problem: https://leetcode.com/problems/binary-tree-level-order-traversal/
 * Author: Araceae ; Date: 2021/4/20
 * Solution: DFS
 * Time Complexity: O(MN)
 * Space Complexity: O(1)
 * Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Level Order Traversal.
 * Memory Usage: 3 MB, less than 9.26% of Go online submissions for Binary Tree Level Order Traversal.
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Problem: https://leetcode.com/problems/binary-tree-level-order-traversal/
func LevelOrder(root *unittest.TreeNode) [][]int {
	res := [][]int{}
	traverseDfs(root, 0, &res)
	return res
}

func traverseDfs(root *unittest.TreeNode, level int, res *[][]int) {
	if root == nil {
		return
	}
	if len(*res) <= level {
		*res = append(*res, []int{root.Val})
	} else {
		(*res)[level] = append((*res)[level], root.Val)
	}
	traverseDfs(root.Left, level+1, res)
	traverseDfs(root.Right, level+1, res)
}
