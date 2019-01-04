package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
		return head
	}
	len, list, headCopy := 0, make([]*ListNode, 0), head
	for headCopy != nil {
		len++
		list = append(list, headCopy)
		headCopy = headCopy.Next
	}
	k = k % len
	if k == 0 {
		return head
	}
	list[len-k-1].Next = nil
	list[len-1].Next = list[0]
	return list[len-k]
}
func main() {

}
