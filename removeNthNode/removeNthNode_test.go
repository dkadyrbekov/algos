package removeNthNode

import (
	"testing"
)

func Test_removeNthLastNode(t *testing.T) {
	type args struct {
		head *ListNode
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
				head: nil,
				n:    0,
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeNthLastNode(tt.args.head, tt.args.n)
			if !equalLists(got, tt.want) {
				t.Errorf("removeNthLastNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equalLists(got *ListNode, wanted []int) bool {

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
