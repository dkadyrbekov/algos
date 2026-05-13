package dynamic

import (
	"reflect"
	"testing"
)

func Test_combinationSum(t *testing.T) {

	tests := []struct {
		name   string
		nums   []int
		target int
		want   [][]int
	}{
		{
			name:   "one",
			nums:   []int{2, 3, 9},
			target: 12,
			want:   [][]int{{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum(tt.nums, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum(%v, %v) = %v, want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
