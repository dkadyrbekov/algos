package greedy

import "testing"

func Test_rescueBoats(t *testing.T) {
	tests := []struct {
		people []int
		limit  int
		want   int
	}{
		{
			people: []int{3, 1, 4, 2, 4},
			limit:  4,
			want:   4,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := rescueBoats(tt.people, tt.limit); got != tt.want {
				t.Errorf("rescueBoats(%v, %v) = %v, want %v", tt.people, tt.limit, got, tt.want)
			}
		})
	}
}

//func Test_binSearch(t *testing.T) {
//
//	tests := []struct {
//		name      string
//		people    []int
//		limit     int
//		wantIndex int
//		wantFound bool
//	}{
//		{
//			name:   "",
//			people: []int{},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			gotIndex, gotFound := binSearch(tt.people, tt.limit)
//			if gotIndex != tt.wantIndex {
//				t.Errorf("binSearch() gotIndex = %v, want %v", gotIndex, tt.wantIndex)
//			}
//			if gotFound != tt.wantFound {
//				t.Errorf("binSearch() gotFound = %v, want %v", gotFound, tt.wantFound)
//			}
//		})
//	}
//}
