package main

/*

给定一个单链表的头节点 head,实现一个调整单链表的函数，使得每K个节点之间为一组进行逆序，
并且从链表的尾部开始组起，头部剩余节点数量不够一组的不需要逆序。（不能使用队列或者栈作为辅助）
例如：
链表:1->2->3->4->5->6->7->8->null, K = 3。那么 6->7->8，3->4->5，1->2各位一组。
调整后：1->2->5->4->3->8->7->6->null。其中 1，2不调整，因为不够一组。
*/

type Node struct {
	Val  int
	Next *Node
}

func batchReverse(head *Node) *Node {
	tail, len := head, 0
	for tail != nil {
		tail = tail.Next
		len++
	}
	times, before := len/3, len%3
	res, prev := head, head
	for i := 0; i < before; i++ {
		prev = head
		head = head.Next

	}
	for i := 0; i < times; i++ {
		if before == 0 && i == 0 {
			res = head.Next.Next
		}
		nextT := head.Next.Next.Next
		head.Next.Next = nil
		prev.Next = reverse(head)
		head.Next = nextT
		head = nextT
	}
	return res
}

func reverse(head *Node) *Node {
	var prev, next *Node
	for head != nil {
		next = head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

func main() {

}
