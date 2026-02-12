package threesum

import (
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {

	tests := []struct {
		name string
		arg  []int
		want [][]int
	}{
		{
			name: "empy list",
			arg:  []int{},
			want: [][]int{},
		},
		{
			name: "one list",
			arg:  []int{1, 0, -1},
			want: [][]int{
				{1, 0, -1},
			},
		},
		{
			name: "list of two",
			arg:  []int{1, -1},
			want: [][]int{},
		},
		{
			name: "no output",
			arg:  []int{1, 1, 9, 15, -3},
			want: [][]int{},
		},
		{
			name: "three output",
			arg:  []int{15, 9, 10, -1, 1, 0, -24, -9},
			want: [][]int{
				{-1, 1, 0},
				{15, 9, -24},
				{10, -1, -9},
				{0, 9, -9},
			},
		},
		{
			name: "resolve unique",
			arg:  []int{15, 9, 10, -1, 1, 0, -24, -9, -9, -1, 1, 0},
			want: [][]int{
				{-1, 1, 0},
				{15, 9, -24},
				{10, -1, -9},
			},
		},
		{
			name: "regular case",
			arg:  []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{
				{-1, 1, 0},
				{-1, -1, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
