package leetcode

// Problem: https://leetcode.com/problems/sort-list/
// Author: Araceae
// Date: 2021/4/24

// Solution: Merge Sort

// Time Complexity: O(NlogN)
// Space Complexity: O(1)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Problem: https://leetcode.com/problems/sort-list/
func SortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := FindMid(head)
	left := SortList(head)
	right := SortList(mid)

	return MergeSort(left, right)
}

func FindMid(head *ListNode) *ListNode {
	slow := head
	fast := head
	var prev *ListNode
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = nil
	return slow
}

func MergeSort(left, right *ListNode) *ListNode {
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
