package leetcode

import (
	"reflect"
	"testing"

	"github.com/araceae101/LeetCode/pkg/unittest"
)

func TestSortList(t *testing.T) {
	type args struct {
		head []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "Example-1", args: args{head: []int{4, 2, 1, 3}}, want: []int{1, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := unittest.ListNode2Slice(SortList(unittest.Slice2ListNode(tt.args.head))); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortList() = %v, want %v", got, tt.want)
			}
		})
	}
}
