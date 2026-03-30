package list_cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) bool {

	fast := head
	slow := head

	for {
		if slow == nil || fast == nil {
			return false
		}

		slow = slow.Next
		fast = fast.Next

		if fast == nil {
			return false
		}

		fast = fast.Next

		if fast == slow {
			return true
		}
	}
}
