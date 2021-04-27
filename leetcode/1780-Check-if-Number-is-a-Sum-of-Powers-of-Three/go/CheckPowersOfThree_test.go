package leetcode

import (
	"testing"
)

func TestCheckPowersOfThreeIterative(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Example-1", args: args{n: 12}, want: true},
		{name: "Example-2", args: args{n: 91}, want: true},
		{name: "MyTest-3", args: args{n: 21}, want: false},
		{name: "MyTest-4", args: args{n: 1}, want: true},
		{name: "MyTest-5", args: args{n: 2}, want: false},
		{name: "MyTest-6", args: args{n: 10000000}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckPowersOfThreeIterative(tt.args.n); got != tt.want {
				t.Errorf("CheckPowersOfThreeIterative() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckPowersOfThreeRecursive(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Example-1", args: args{n: 12}, want: true},
		{name: "Example-2", args: args{n: 91}, want: true},
		{name: "MyTest-3", args: args{n: 21}, want: false},
		{name: "MyTest-4", args: args{n: 1}, want: true},
		{name: "MyTest-5", args: args{n: 2}, want: false},
		{name: "MyTest-6", args: args{n: 10000000}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckPowersOfThreeRecursive(tt.args.n); got != tt.want {
				t.Errorf("CheckPowersOfThreeRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
