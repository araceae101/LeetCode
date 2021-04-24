package unittest

type ListNode struct {
	Val  int
	Next *ListNode
}

// Convert []int into *ListNode
func Slice2ListNode(nums []int) *ListNode {
	l := len(nums)

	if l == 0 {
		return nil
	}

	res := &ListNode{}
	curr := res

	for i := 0; i < l; i++ {
		curr.Next = &ListNode{
			Val: nums[i],
		}
		curr = curr.Next
	}

	return res.Next
}
