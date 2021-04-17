package leetcode

import (
	"reflect"
	"testing"
)

func TestUpdateMatrixBFS(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpdateMatrix(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateMatrixBFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
