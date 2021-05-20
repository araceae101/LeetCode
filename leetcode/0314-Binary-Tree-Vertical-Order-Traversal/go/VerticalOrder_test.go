package leetcode

import (
	"reflect"
	"testing"

	"github.com/araceae101/LeetCode/pkg/unittest"
)

func TestVerticalOrder(t *testing.T) {
	type args struct {
		root *unittest.TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "Example-1", args: args{unittest.Slice2TreeNode([]int{3, 9, 20, -1, -1, 15, 7})}, want: [][]int{{9}, {3, 15}, {20}, {7}}},
		{name: "Example-2", args: args{unittest.Slice2TreeNode([]int{3, 9, 8, 4, 0, 1, 7})}, want: [][]int{{4}, {9}, {3, 0, 1}, {8}, {7}}},
		{name: "Example-3", args: args{unittest.Slice2TreeNode([]int{3, 9, 8, 4, 0, 1, 7, -1, -1, -1, 2, 5})}, want: [][]int{{4}, {9, 5}, {3, 0, 1}, {8, 2}, {7}}},
		{name: "Example-4", args: args{unittest.Slice2TreeNode([]int{})}, want: [][]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VerticalOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VerticalOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
