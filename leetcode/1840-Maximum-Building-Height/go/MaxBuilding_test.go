package leetcode

import "testing"

func TestMaxBuilding(t *testing.T) {
	type args struct {
		n int
		A [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Example-1", args: args{n: 5, A: [][]int{{2, 1}, {4, 1}}}, want: 2},
		{name: "Example-2", args: args{n: 6, A: [][]int{}}, want: 5},
		{name: "Example-3", args: args{n: 10, A: [][]int{{5, 3}, {2, 5}, {7, 4}, {10, 3}}}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxBuilding(tt.args.n, tt.args.A); got != tt.want {
				t.Errorf("MaxBuilding() = %v, want %v", got, tt.want)
			}
		})
	}
}
