package unittest

// TreeNode is structure for Binary Tree
// positive integer in Val represents the valid node and -1 represents null node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Slice2TreeNode Converts []int int into *TreeNode
func Slice2TreeNode(nums []int) *TreeNode {
	l := len(nums)
	if l == 0 {
		return nil
	}
	root := new(TreeNode)
	root.Val = nums[0]

	ch := make(chan *TreeNode, len(nums))
	ch <- root
	nums = nums[1:]

	for i := 0; i < len(nums); i++ {
		tree := <-ch
		// build left tree
		if i < len(nums) && nums[i] == -1 {
			tree.Left = nil
		} else if i < len(nums) {
			tree.Left = &TreeNode{
				Val: nums[i],
			}
			ch <- tree.Left
		}
		i++
		// build right tree
		if i < len(nums) && nums[i] == -1 {
			tree.Right = nil
		} else if i < len(nums) {
			tree.Right = &TreeNode{
				Val: nums[i],
			}
			ch <- tree.Right
		}
	}
	return root
}
