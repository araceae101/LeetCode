package leetcode

import "testing"

func TestMinTime(t *testing.T) {
	type args struct {
		n        int
		edges    [][]int
		hasApple []bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Exampe-1",
			args: args{n: 7, edges: [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}, hasApple: []bool{false, false, true, false, true, true, false}},
			want: 8,
		},
		{
			name: "Exampe-2",
			args: args{n: 7, edges: [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}, hasApple: []bool{false, false, true, false, false, true, false}},
			want: 6,
		},
		{
			name: "Exampe-3",
			args: args{n: 7, edges: [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}, hasApple: []bool{false, false, false, false, false, false, false}},
			want: 0,
		},
		{
			name: "MyTest-1",
			args: args{n: 5, edges: [][]int{{0, 1}, {0, 2}, {1, 3}, {0, 4}}, hasApple: []bool{false, false, false, true, false}},
			want: 4,
		},
		{
			name: "MyTest-2",
			args: args{n: 4, edges: [][]int{{0, 2}, {0, 3}, {1, 2}}, hasApple: []bool{false, true, false, false}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinTimeBFS(tt.args.n, tt.args.edges, tt.args.hasApple); got != tt.want {
				t.Errorf("MinTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
