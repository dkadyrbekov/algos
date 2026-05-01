package backtracking

import (
	"reflect"
	"testing"
)

func Test_restoreIpAddresses(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []string
	}{
		{
			name: "one",
			s:    "010010",
			want: []string{"0.10.0.10", "0.100.1.0"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := restoreIpAddresses(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoreIpAddresses(%v) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
