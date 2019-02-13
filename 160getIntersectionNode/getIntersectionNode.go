package _60getIntersectionNode

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	all := make(map[*ListNode]*struct{})
	for headA != nil {
		all[headA] = &struct{}{}
		headA = headA.Next
	}
	for headB != nil {
		if all[headB] != nil {
			return headB
		}
		headB = headB.Next
	}
	return nil
}
