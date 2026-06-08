package dynamic

import "testing"

func Test_countGoodSubsequences(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "empty",
			s:    "",
			want: 0,
		},
		{
			name: "aa",
			s:    "aa",
			want: 3,
		},
		{
			name: "good",
			s:    "good",
			want: 12,
		},
		{
			name: "goodd",
			s:    "good",
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countGoodSubsequences(tt.s); got != tt.want {
				t.Errorf("countGoodSubsequences(%v) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}

func Test_isGoodSubsequence(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "empty",
			s:    "",
			want: true,
		},
		{
			name: "aa",
			s:    "aa",
			want: true,
		},
		{
			name: "good",
			s:    "good",
			want: false,
		},
		{
			name: "goodd",
			s:    "goodd",
			want: false,
		},
		{
			name: "ggoodd",
			s:    "ggoodd",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isGoodSubsequence(tt.s); got != tt.want {
				t.Errorf("isGoodSubsequence(%v) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
