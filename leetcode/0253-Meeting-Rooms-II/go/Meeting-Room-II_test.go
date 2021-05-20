package leetcode

import "testing"

func TestMinMeetingRooms(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Example-1", args: args{[][]int{{0, 30}, {5, 10}, {15, 20}}}, want: 2},
		{name: "Example-2", args: args{[][]int{{7, 10}, {2, 4}}}, want: 1},
		{name: "Mytest-1", args: args{[][]int{{7, 10}}}, want: 1},
		{name: "Mytest-2", args: args{[][]int{}}, want: 0},
		{name: "Mytest-3", args: args{[][]int{{0, 5}, {5, 10}, {10, 15}}}, want: 1},
		{name: "Mytest-4", args: args{[][]int{{0, 5}, {10, 15}, {5, 10}}}, want: 1},
		{name: "Mytest-5", args: args{[][]int{{0, 5}, {0, 5}, {0, 5}}}, want: 3},
		{name: "Mytest-6", args: args{[][]int{{0, 10}, {5, 20}, {15, 30}}}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinMeetingRooms(tt.args.intervals); got != tt.want {
				t.Errorf("MinMeetingRooms() = %v, want %v", got, tt.want)
			}
		})
	}
}
