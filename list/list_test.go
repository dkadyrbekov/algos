package list

import (
	"testing"
)

func Test_reorderList(t *testing.T) {
	tests := []struct {
		name string
		list []int
	}{
		{
			name: "list",
			list: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name: "list",
			list: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reorderList(makeList(tt.list))
		})
	}
}

func Test_reverseBetween(t *testing.T) {
	tests := []struct {
		name  string
		list  []int
		left  int
		right int
	}{
		{
			name:  "list",
			list:  []int{1, 2, 3, 4, 5, 6, 7, 8},
			left:  1,
			right: 4,
		},
		{
			name:  "list",
			list:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			left:  3,
			right: 5,
		},
		{
			name:  "list",
			list:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			left:  3,
			right: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseBetween(makeList(tt.list), tt.left, tt.right)
		})
	}
}

func Test_reverseEvenGroup(t *testing.T) {
	tests := []struct {
		name string
		list []int
	}{
		{
			name: "list",
			list: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name: "list",
			list: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "list",
			list: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseEvenLengthGroups(makeList(tt.list))
		})
	}
}

func makeList(values []int) *ListNode {
	head := &ListNode{
		Val:  values[0],
		Next: nil,
	}

	curr := head
	for i := 1; i < len(values); i++ {
		curr.Next = &ListNode{
			Val:  values[i],
			Next: nil,
		}

		curr = curr.Next
	}

	return head
}
