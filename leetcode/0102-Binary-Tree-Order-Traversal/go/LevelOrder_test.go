package leetcode

import (
	"reflect"
	"testing"

	"github.com/araceae101/LeetCode/pkg/unittest"
)

func TestLevelOrder(t *testing.T) {
	type args struct {
		root *unittest.TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "Example-1", args: args{unittest.Slice2TresNode([]int{3, 9, 20, -1, -1, 15, 7})}, want: [][]int{{3}, {9, 20}, {15, 7}}},
		{name: "Example-2", args: args{unittest.Slice2TresNode([]int{1})}, want: [][]int{{1}}},
		{name: "Example-3", args: args{unittest.Slice2TresNode([]int{})}, want: [][]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LevelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LevelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
