package leetcode

import t "github.com/araceae101/LeetCode/pkg/unittest"

// Problem: https://leetcode.com/problems/sort-list/
// Author: Araceae
// Date: 2021/4/24

// Solution: Merge Sort

// Time Complexity: O(NlogN)
// Space Complexity: O(1)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Problem: https://leetcode.com/problems/sort-list/
func SortList(head *t.ListNode) *t.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := FindMid(head)
	left := SortList(head)
	right := SortList(mid)

	return MergeSort(left, right)
}

func FindMid(head *t.ListNode) *t.ListNode {
	slow := head
	fast := head
	var prev *t.ListNode
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = nil
	return slow
}

func MergeSort(left, right *t.ListNode) *t.ListNode {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	if left.Val <= right.Val {
		left.Next = MergeSort(left.Next, right)
		return left
	}
	right.Next = MergeSort(left, right.Next)
	return right
}
