package subsets

import (
	"reflect"
	"testing"
)

func Test_letterCombinations(t *testing.T) {
	tests := []struct {
		name   string
		digits string
		want   []string
	}{
		{
			name:   "23",
			digits: "23",
			want:   []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := letterCombinations(tt.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("letterCombinations(%v) = %v, want %v", tt.digits, got, tt.want)
			}
		})
	}
}
