package interview45_Last_Ramining_Num

type Node struct {
	Val  int
	Next *Node
}

func LastRaminNum(n int) int{
	head := &Node{Val: 0}
	current := head
	for i := 1; i <= n; i++ {
		if i != n {
			current.Next = &Node{Val: i}
			current = current.Next
		} else {
			current.Next = &Node{Val: i, Next: head}
		}
	}
	for head.Next != head {
		head.Next.Next = head.Next.Next.Next
		head = head.Next.Next
	}
	return head.Val
}
