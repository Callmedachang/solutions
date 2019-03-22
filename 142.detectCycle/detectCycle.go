package _42_detectCycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	// 步骤一：使用快慢指针判断链表是否有环
	p, p2 := head, head
	hasCycle := false
	for p2.Next != nil && p2.Next.Next != nil {
		p = p.Next
		p2 = p2.Next.Next
		if p == p2 {
			hasCycle = true
			break
		}
	}
	// 步骤二：若有环，找到入环开始的节点
	if hasCycle {
		q := head
		for p != q {
			p = p.Next
			q = q.Next
		}
		return q
	}
	return nil
}
