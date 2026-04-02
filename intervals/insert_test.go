package intervals

import (
	"reflect"
	"testing"
)

func Test_insertInterval(t *testing.T) {
	tests := []struct {
		name              string
		existingIntervals [][]int
		newInterval       []int
		want              [][]int
	}{
		{
			name:              "empty existingIntervals 1",
			existingIntervals: [][]int{},
			newInterval:       []int{1, 2},
			want: [][]int{
				{1, 2},
			},
		},
		{
			name: "empty existingIntervals 2",
			existingIntervals: [][]int{
				{1, 4}, {5, 8}, {11, 17},
			},
			newInterval: []int{9, 10},
			want: [][]int{
				{1, 4}, {5, 8}, {9, 10}, {11, 17},
			},
		},
		{
			name: "empty existingIntervals 3",
			existingIntervals: [][]int{
				{1, 4}, {5, 8}, {11, 17},
			},
			newInterval: []int{8, 12},
			want: [][]int{
				{1, 4}, {5, 17},
			},
		},
		{
			name: "empty existingIntervals 4",
			existingIntervals: [][]int{
				{1, 4}, {5, 8}, {11, 17},
			},
			newInterval: []int{1, 20},
			want: [][]int{
				{1, 20},
			},
		},
		{
			name: "empty existingIntervals 5",
			existingIntervals: [][]int{
				{1, 4}, {5, 8}, {11, 17},
			},
			newInterval: []int{20, 21},
			want: [][]int{
				{1, 4}, {5, 8}, {11, 17}, {20, 21},
			},
		},
		{
			name: "empty existingIntervals 6",
			existingIntervals: [][]int{
				{5, 8}, {11, 17},
			},
			newInterval: []int{1, 2},
			want: [][]int{
				{1, 2}, {5, 8}, {11, 17},
			},
		},
		{
			name: "empty existingIntervals 7",
			existingIntervals: [][]int{
				{5, 8}, {11, 17},
			},
			newInterval: []int{1, 6},
			want: [][]int{
				{1, 8}, {11, 17},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newIntArg := make([]int, len(tt.newInterval))
			copy(newIntArg, tt.newInterval)
			if got := insertInterval(tt.existingIntervals, tt.newInterval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertInterval(%v, %v) = %v, want %v", tt.existingIntervals, newIntArg, got, tt.want)
			}
		})
	}
}
