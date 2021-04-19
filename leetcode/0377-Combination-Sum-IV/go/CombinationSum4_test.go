package leetdocde

import "testing"

func TestCombinationSum4(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Example-1", args: args{nums: []int{1, 2, 3}, target: 4}, want: 7},
		{name: "Example-2", args: args{nums: []int{9}, target: 3}, want: 0},
		{name: "MyTest-1", args: args{nums: []int{1, 2, 3}, target: 32}, want: 181997601},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CombinationSum4(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("CombinationSum4() = %v, want %v", got, tt.want)
			}
		})
	}
}
