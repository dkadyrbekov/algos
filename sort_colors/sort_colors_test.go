package sort_colors

import (
	"reflect"
	"testing"
)

func Test_sortColors(t *testing.T) {
	tests := []struct {
		name   string
		colors []int
		want   []int
	}{
		{
			name:   "empty",
			colors: []int{},
			want:   []int{},
		},
		{
			name:   "invalid color in input array",
			colors: []int{1, 0, 2, 3, -1, 0},
			want:   []int{},
		},
		{
			name:   "normal array",
			colors: []int{1, 0, 2, 2, 0, 0},
			want:   []int{0, 0, 0, 1, 2, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortColors(tt.colors); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortColors(%v) = %v, want %v", tt.colors, got, tt.want)
			}
		})
	}
}
