package repeating_char

import "testing"

func Test_longestRepeatingCharacterReplacement(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{
			name: "empty string with k = 0",
			s:    "",
			k:    0,
			want: 0,
		},
		{
			name: "empty string with k = 1",
			s:    "",
			k:    1,
			want: 0,
		},
		{
			name: "valid string",
			s:    "AABCCBB",
			k:    2,
			want: 5,
		},
		{
			name: "valid string",
			s:    "COOLLOOC",
			k:    2,
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestRepeatingCharacterReplacement(tt.s, tt.k); got != tt.want {
				t.Errorf("longestRepeatingCharacterReplacement(%v, %v) = %v, want %v", tt.s, tt.k, got, tt.want)
			}
		})
	}
}

//func Test_repeatingCond(t *testing.T) {
//	tests := []struct {
//		name string
//		s    []rune
//		k    int
//		want bool
//	}{
//		{
//			name: "empty string with k = 0",
//			s:    []rune(""),
//			k:    0,
//			want: true,
//		},
//		{
//			name: "empty string with k = 1",
//			s:    []rune(""),
//			k:    1,
//			want: true,
//		},
//		{
//			name: "String with enough k",
//			s:    []rune("AABCC"),
//			k:    2,
//			want: false,
//		},
//		{
//			name: "String with not enough k",
//			s:    []rune("AABCC"),
//			k:    3,
//			want: true,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := repeatingCond(tt.s, tt.k); got != tt.want {
//				t.Errorf("repeatingCond(%v, %v) = %v, want %v", tt.s, tt.k, got, tt.want)
//			}
//		})
//	}
//}
