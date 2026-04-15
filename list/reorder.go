package list

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	slow, fast := head, head
	var head2 *ListNode
	for {
		if fast.Next == nil || fast.Next.Next == nil {
			head2 = slow.Next
			slow.Next = nil
			break
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	var prev *ListNode
	current := head2
	next := head2.Next
	for {
		current.Next = prev
		prev = current
		if next == nil {
			break
		}
		current = next
		next = next.Next
	}

	head1 := head
	head2 = current
	current = head1

	for head1 != nil && head2 != nil {
		head1 = head1.Next

		current.Next = head2
		current = head2

		head2 = head2.Next
		current.Next = head1
		current = head1
	}
}

func printList(l *ListNode) string {
	var s strings.Builder

	for l != nil {
		s.WriteString(fmt.Sprintf("%v ->", l.Val))
		l = l.Next
	}

	s.WriteString(fmt.Sprintf("\n"))

	return s.String()
}
