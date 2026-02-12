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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unique(t *testing.T) {

	tests := []struct {
		name     string
		zeroSums [][]int
		want     [][]int
	}{
		{
			name:     "empty",
			zeroSums: [][]int{},
			want:     [][]int{},
		},
		{
			name: "one list",
			zeroSums: [][]int{
				{1, 2, 3},
			},
			want: [][]int{
				{1, 2, 3},
			},
		},
		{
			name: "resolve unique",
			zeroSums: [][]int{
				{1, 2, 3},
				{5, 4, 3},
				{2, 1, 3},
				{2, 1, 7},
			},
			want: [][]int{
				{1, 2, 3},
				{5, 4, 3},
				{1, 2, 7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := unique(tt.zeroSums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unique() = %v, want %v", got, tt.want)
			}
		})
	}
}
