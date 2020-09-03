package _2reverse2

/*
反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。

说明:
1 ≤ m ≤ n ≤ 链表长度。

示例:

输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseBetween(head *ListNode, m int, n int) *ListNode {
	if m >= n || head == nil || head.Next == nil {
		return head
	}
	var pre1, pre2 *ListNode
	var mNode, nNode *ListNode
	cur := head
	for l := 1; l < n; l++ {
		if l < m {
			pre1 = cur
		}
		if l >= m {
			pre2 = cur
		}
		if cur == nil {
			return head
		}
		cur = cur.Next
	}
	if pre1 == nil {
		mNode, nNode = head, pre2.Next
	} else {
		mNode, nNode = pre1.Next, pre2.Next
	}
	n3 := nNode.Next
	if pre1 != nil {
		pre1.Next = nil
	}
	nNode.Next = nil
	temp := mNode
	rn2 := r(mNode)
	if pre1!=nil{
		pre1.Next = rn2
	}else{
		head = rn2
	}
	if temp != nil {
		temp.Next = n3
	}
	return head
}

func r(head *ListNode) *ListNode {
	var pre, next *ListNode
	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}
