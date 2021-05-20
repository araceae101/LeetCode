package leetcode

import "github.com/araceae101/LeetCode/pkg/unittest"

/**
 * Problem: https://leetcode.com/problems/binary-tree-vertical-order-traversal/
 * Author: Araceae ; Date: 2021/4/20
 * Solution: BFS (use channel and hash map)
 * Time Complexity: O(N)
 * Space Complexity: O(N)
 * Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Vertical Order Traversal.
 * Memory Usage: 2.6 MB, less than 22.86% of Go online submissions for Binary Tree Vertical Order Traversal.
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type myNode struct {
	Node *unittest.TreeNode
	Idx  int
}

// Problem: https://leetcode.com/problems/binary-tree-vertical-order-traversal/
func VerticalOrder(root *unittest.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	r := new(myNode)
	r.Node = root
	r.Idx = 0

	ch := make(chan *myNode, 100)
	ch <- r

	m := make(map[int][]int)

	for len(ch) > 0 {
		currNode := <-ch
		_, ok := m[currNode.Idx]
		if !ok {
			m[currNode.Idx] = []int{currNode.Node.Val}
		} else {
			m[currNode.Idx] = append(m[currNode.Idx], currNode.Node.Val)
		}

		if currNode.Node.Left != nil {
			ch <- &myNode{
				Node: currNode.Node.Left,
				Idx:  currNode.Idx - 1,
			}
		}
		if currNode.Node.Right != nil {
			ch <- &myNode{
				Node: currNode.Node.Right,
				Idx:  currNode.Idx + 1,
			}
		}
	}

	res := make([][]int, 0)
	for i := -100; i < 101; i++ {
		if _, ok := m[i]; ok {
			res = append(res, m[i])
		}
	}
	return res
}
