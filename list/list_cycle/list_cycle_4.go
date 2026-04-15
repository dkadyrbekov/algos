package list_cycle

func removeCycle(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	fast, slow := head, head

	for {
		slow = slow.Next
		if fast == nil || fast.Next == nil {
			return head
		}
		fast = fast.Next.Next

		if slow == fast {
			break
		}
	}

	slow = head

	for {
		if fast == slow {
			break
		}

		fast = fast.Next
		slow = slow.Next
	}

	for {
		if fast.Next == slow {
			fast.Next = nil
			return head
		}

		fast = fast.Next
	}
}
