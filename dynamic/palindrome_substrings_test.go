package dynamic

import "testing"

func Test_isPalindrome(t *testing.T) {

	tests := []struct {
		s    string
		want bool
	}{
		{
			s:    "a",
			want: true,
		},
		{
			s:    "ala",
			want: true,
		},
		{
			s:    "olololo",
			want: true,
		},
		{
			s:    "ol",
			want: false,
		},
		{
			s:    "olololol",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := isPalindrome(tt.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
