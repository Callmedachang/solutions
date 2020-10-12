package _5reverseKGroup

/*
给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。

k 是一个正整数，它的值小于或等于链表的长度。

如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

 

示例：

给你这个链表：1->2->3->4->5

当 k = 2 时，应当返回: 2->1->4->3->5

当 k = 3 时，应当返回: 3->2->1->4->5

 

说明：

你的算法只能使用常数的额外空间。
你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseKGroup(head *ListNode, k int) *ListNode {
	var trueHead, prev *ListNode
	before := head
	end, needReverse := getKNext(head, k)
	for before != nil {
		//当前需要反转的最后一个node
		temp := end.Next
		//断开最后一个node的Next
		end.Next = nil
		//翻转拿到head
		var newBefore *ListNode
		if needReverse{
			newBefore = reverse(before)
		}else{
			newBefore = before
		}

		if prev != nil {
			//如果之前链路有数据那么之前的最后一个节点的Next等于现在的head
			prev.Next = newBefore
		} else {
			trueHead = newBefore
		}
		//拿到最后一个节点
		for newBefore.Next != nil {
			newBefore = newBefore.Next
		}
		prev = newBefore
		before = temp
		end, needReverse = getKNext(temp, k)
	}
	return trueHead
}

func getKNext(head *ListNode, k int) (*ListNode, bool) {
	if head == nil {
		return nil, false
	}
	for k > 1 {
		if head.Next == nil {
			return head, false
		}
		head = head.Next
		k--
	}
	return head, true
}

func reverse(head *ListNode) *ListNode {
	var prev, next *ListNode
	for head != nil {
		next = head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}
