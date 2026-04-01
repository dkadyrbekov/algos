package intervals

import (
	"reflect"
	"testing"
)

func Test_mergeIntervals(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      [][]int
	}{
		{
			name:      "empty",
			intervals: [][]int{},
			want:      [][]int{},
		},
		{
			name: "valid",
			intervals: [][]int{
				{7, 10}, {8, 9}, {1, 5}, {6, 7},
			},
			want: [][]int{
				{1, 5}, {6, 10},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeIntervals(tt.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeIntervals(%v) = %v, want %v", tt.intervals, got, tt.want)
			}
		})
	}
}
