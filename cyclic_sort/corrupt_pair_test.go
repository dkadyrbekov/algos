package cyclic_sort

import (
	"reflect"
	"testing"
)

func Test_findCorruptPair(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "empty",
			nums: []int{},
			want: []int{},
		},
		{
			name: "normal",
			nums: []int{7, 1, 2, 3, 1, 5, 6},
			want: []int{4, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCorruptPair(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findCorruptPair(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
