package leetcode

import (
	"reflect"
	"testing"
)

func TestSortList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "Example-1", args: args{head: Slice2ListNode([]int{4, 2, 1, 3})}, want: Slice2ListNode([]int{4, 2, 1, 3})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortList() = %v, want %v", got, tt.want)
			}
		})
	}
}
