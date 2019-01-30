package _41HasCycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	all := make(map[*ListNode]*struct{})
	for head != nil {
		if all[head] == nil {
			all[head] = &struct{}{}
		} else {
			return true
		}
		head = head.Next
	}
	return false
}
