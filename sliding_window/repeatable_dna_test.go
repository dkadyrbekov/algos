package sliding_window

import (
	"testing"
)

func Test_findRepeatedDnaSequences(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []string
	}{
		{
			name: "empty string",
			s:    "",
			want: []string{},
		},
		{
			name: "valid string",
			s:    "aaaaacccccaaaaacccccaaaaat",
			want: []string{"aacccccaaa", "aaacccccaa", "acccccaaaa", "cccccaaaaa", "aaaaaccccc", "aaaaccccca"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRepeatedDnaSequences(tt.s); !equalStringSlices(got, tt.want) {
				t.Errorf("findRepeatedDnaSequences(%v) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}

func equalStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	count := make(map[string]int)

	for _, s := range a {
		count[s]++
	}

	for _, s := range b {
		if count[s] == 0 {
			return false
		}
		count[s]--
	}

	return true
}
