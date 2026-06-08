package dynamic

import "testing"

func Test_numOfDecodings(t *testing.T) {

	tests := []struct {
		decodeStr string
		want      int
	}{
		{
			decodeStr: "2",
			want:      1,
		},
		{
			decodeStr: "22",
			want:      2,
		},
		{
			decodeStr: "222",
			want:      3,
		},
		{
			decodeStr: "2222",
			want:      5,
		},
		{
			decodeStr: "22222",
			want:      8,
		},
		{
			decodeStr: "222222",
			want:      13,
		},
		{
			decodeStr: "2222222",
			want:      21,
		},
		{
			decodeStr: "22222222",
			want:      34,
		},
		{
			decodeStr: "32535223",
			want:      34,
		},
	}
	for _, tt := range tests {
		t.Run(tt.decodeStr, func(t *testing.T) {
			if got := numOfDecodings(tt.decodeStr); got != tt.want {
				t.Errorf("numOfDecodings() = %v, want %v", got, tt.want)
			}
		})
	}
}
