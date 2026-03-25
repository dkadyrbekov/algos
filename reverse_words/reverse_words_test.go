package reverse_words

import "testing"

func Test_reverseWords(t *testing.T) {
	tests := []struct {
		name     string
		sentence string
		want     string
	}{
		{
			name:     "empty",
			sentence: "",
			want:     "",
		},
		{
			name:     "only spaces",
			sentence: "       ",
			want:     "",
		},
		{
			name:     "one word",
			sentence: "word",
			want:     "word",
		},
		{
			name:     "sentence without extra spaces",
			sentence: "three word spaces",
			want:     "spaces word three",
		},
		{
			name:     "lot of spaces",
			sentence: "   three   word    spaces    ",
			want:     "spaces word three",
		},
		{
			name:     "one letter",
			sentence: "e",
			want:     "e",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.sentence); got != tt.want {
				t.Errorf("reverseWords(\"%v\") = \"%v\", want \"%v\"", tt.sentence, got, tt.want)
			}
		})
	}
}
