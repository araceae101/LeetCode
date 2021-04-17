package leetcode

import (
	"reflect"
	"testing"
)

type test struct {
	name string
	args [][]int
	want [][]int
}

var testcases = []test{
	{name: "example-1", args: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, want: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}},
	{name: "example-2", args: [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}, want: [][]int{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}}},
	{name: "mytest-1", args: [][]int{{}}, want: [][]int{{}}},
	{name: "mytest-2", args: [][]int{{0, 0, 0, 0, 0, 0}}, want: [][]int{{0, 0, 0, 0, 0, 0}}},
	{name: "mytest-3", args: [][]int{{1, 0, 0, 0, 0, 0}}, want: [][]int{{1, 0, 0, 0, 0, 0}}},
	{name: "mytest-4", args: [][]int{{0, 0, 0, 0, 0, 1}}, want: [][]int{{0, 0, 0, 0, 0, 1}}},
	{name: "mytest-5", args: [][]int{{0, 0, 0, 1, 0, 0}}, want: [][]int{{0, 0, 0, 1, 0, 0}}},
	{name: "mytest-6", args: [][]int{{0}, {0}, {0}, {0}, {0}}, want: [][]int{{0}, {0}, {0}, {0}, {0}}},
	{name: "mytest-7", args: [][]int{{1}, {0}, {0}, {0}, {0}}, want: [][]int{{1}, {0}, {0}, {0}, {0}}},
	{name: "mytest-8", args: [][]int{{0}, {0}, {0}, {0}, {1}}, want: [][]int{{0}, {0}, {0}, {0}, {1}}},
	{name: "mytest-9", args: [][]int{{0}, {0}, {1}, {0}, {0}}, want: [][]int{{0}, {0}, {1}, {0}, {0}}},
	{name: "mytest-10", args: [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}, want: [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}},
	{name: "mytest-11", args: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, want: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}},
	{name: "mytest-12", args: [][]int{{0, 1, 1}, {1, 1, 1}, {1, 1, 1}}, want: [][]int{{0, 1, 2}, {1, 2, 3}, {2, 3, 4}}},
	{name: "mytest-13", args: [][]int{{1, 0, 1}, {1, 1, 1}, {1, 1, 1}}, want: [][]int{{1, 0, 1}, {2, 1, 2}, {3, 2, 3}}},
	{name: "mytest-14", args: [][]int{{1, 1, 0}, {1, 1, 1}, {1, 1, 1}}, want: [][]int{{2, 1, 0}, {3, 2, 1}, {4, 3, 2}}},
	{name: "mytest-15", args: [][]int{{1, 1, 1}, {0, 1, 1}, {1, 1, 1}}, want: [][]int{{1, 2, 3}, {0, 1, 2}, {1, 2, 3}}},
	{name: "mytest-16", args: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, want: [][]int{{2, 1, 2}, {1, 0, 1}, {2, 1, 2}}},
	{name: "mytest-17", args: [][]int{{1, 1, 1}, {1, 1, 0}, {1, 1, 1}}, want: [][]int{{3, 2, 1}, {2, 1, 0}, {3, 2, 1}}},
	{name: "mytest-18", args: [][]int{{1, 1, 1}, {1, 1, 1}, {0, 1, 1}}, want: [][]int{{2, 3, 4}, {1, 2, 3}, {0, 1, 2}}},
	{name: "mytest-19", args: [][]int{{1, 1, 1}, {1, 1, 1}, {1, 0, 1}}, want: [][]int{{3, 2, 3}, {2, 1, 2}, {1, 0, 1}}},
	{name: "mytest-20", args: [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 0}}, want: [][]int{{4, 3, 2}, {3, 2, 1}, {2, 1, 0}}},
}

func TestUpdateMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpdateMatrix(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
