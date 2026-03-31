package happy_number

import "testing"

func Test_countNextNumber(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want int
	}{
		{
			name: "zero",
			num:  0,
			want: 0,
		},
		{
			name: "negative",
			num:  -11,
			want: 0,
		},
		{
			name: "one digit",
			num:  2,
			want: 4,
		},
		{
			name: "regular number",
			num:  258,
			want: 93,
		},
		{
			name: "multiple of 10",
			num:  1000,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNextNumber(tt.num); got != tt.want {
				t.Errorf("countNextNumber(%v) = %v, want %v", tt.num, got, tt.want)
			}
		})
	}
}

func Test_isHappy(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want bool
	}{
		{
			name: "zero",
			num:  0,
			want: false,
		},
		{
			name: "negative",
			num:  -11,
			want: false,
		},
		{
			name: "one digit number true",
			num:  7,
			want: true,
		},
		{
			name: "one digit number false",
			num:  9,
			want: false,
		},
		{
			name: "regular happy number",
			num:  103,
			want: true,
		},
		{
			name: "regular not happy number",
			num:  257,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHappy(tt.num); got != tt.want {
				t.Errorf("isHappy(%v) = %v, want %v", tt.num, got, tt.want)
			}
		})
	}
}
