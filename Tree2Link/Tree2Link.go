package Tree2Link

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

var pre, head *Node = nil, nil

func Tree2Link(n *Node) {
	if n == nil {
		return
	}
	Tree2Link(n.Left)
	n.Left = pre
	if pre != nil {
		pre.Right = n
	}
	pre = n
	if head == nil {
		head = n
	}
	Tree2Link(n.Right)
}
