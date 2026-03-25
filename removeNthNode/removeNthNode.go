package removeNthNode

// Definition for a Linked List node

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthLastNode(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	if n <= 0 {
		return head
	}

	prev := &ListNode{Next: head}
	gap := 0

	ptr := head
	for ptr != nil {
		if gap >= n {
			prev = prev.Next
		}

		ptr = ptr.Next
		gap++
	}

	if prev.Next == head {
		return head.Next
	}

	prev.Next = prev.Next.Next

	return head
}
