package removeNthNode

import (
	"fmt"
	"strings"
	"testing"
)

func Test_removeNthLastNode(t *testing.T) {
	type args struct {
		list []int
		n    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "empty list",
			args: args{
				list: []int{},
				n:    0,
			},
			want: []int{},
		},
		{
			name: "Length One",
			args: args{
				list: []int{1},
				n:    1,
			},
			want: []int{},
		},
		{
			name: "Length Two",
			args: args{
				list: []int{1, 2},
				n:    1,
			},
			want: []int{1},
		},
		{
			name: "Length Three remove last",
			args: args{
				list: []int{1, 2, 3},
				n:    1,
			},
			want: []int{1, 2},
		},
		{
			name: "Length Three remove middle",
			args: args{
				list: []int{1, 2, 3},
				n:    2,
			},
			want: []int{1, 3},
		},
		{
			name: "Length Three remove first",
			args: args{
				list: []int{1, 2, 3},
				n:    3,
			},
			want: []int{2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := makeList(tt.args.list)
			got := removeNthLastNode(list, tt.args.n)
			if !equalLists(got, tt.want) {
				t.Errorf("removeNthLastNode() = %v, want %v", listToString(got), tt.want)
			}
		})
	}
}

func makeList(listArr []int) (list *ListNode) {
	var last *ListNode

	for _, num := range listArr {
		if list == nil || last == nil {
			list = &ListNode{
				Val:  num,
				Next: nil,
			}
			last = list
			continue
		}

		last.Next = &ListNode{
			Val:  num,
			Next: nil,
		}
		last = last.Next
	}

	return
}

func equalLists(got *ListNode, wanted []int) bool {

	// check length is equal
	var listLength int
	for l := got; l != nil; l = l.Next {
		listLength++
	}

	if listLength != len(wanted) {
		return false
	}

	// check values are equal
	for i := 0; i < len(wanted); i++ {
		if got == nil {
			return false
		}

		if wanted[i] != got.Val {
			return false
		}

		got = got.Next
	}

	return true
}

func listToString(head *ListNode) string {
	var sb strings.Builder
	for node := head; node != nil; node = node.Next {
		if sb.Len() > 0 {
			sb.WriteString(" -> ")
		}
		sb.WriteString(fmt.Sprintf("%d", node.Val))
	}
	return sb.String()
}
