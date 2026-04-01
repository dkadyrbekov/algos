package sliding_window

import "testing"

func Test_findLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want int
	}{
		{
			name: "string len = 0",
			str:  "",
			want: 0,
		},
		{
			name: "string len = 1",
			str:  "A",
			want: 1,
		},
		{
			name: "regular string",
			str:  "BACDAAFGKMNO",
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLongestSubstring(tt.str); got != tt.want {
				t.Errorf("findLongestSubstring(%v) = %v, want %v", tt.str, got, tt.want)
			}
		})
	}
}
