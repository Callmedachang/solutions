package interview_26_Complex_Link_List_Copy

/*
请实现函数ComplexListNode*Clone（ComplexListNode*pHead），
复制一个复杂链表。在复杂链表中，每个结点除了有一个m_pNext指针指向下一个结点外，
还有一个m_pSibling指向链表中的任意结点或者NULL
*/

type Node struct {
	Val  int
	Next *Node
	Sib  *Node
}

func ComplexListNodeClone(head *Node) *Node {
	if head == nil {
		return head
	}
	copyHead, sibMap := &Node{Val: head.Val}, make(map[*Node]*Node)
	sibMap[head] = copyHead
	newHeadIndex, headIndex := copyHead, head
	for headIndex.Next != nil {
		newHeadIndex.Next = &Node{Val: headIndex.Next.Val}
		headIndex = headIndex.Next
		sibMap[headIndex] = newHeadIndex.Next
		newHeadIndex = newHeadIndex.Next
	}
	newHeadIndex, headIndex = copyHead, head
	for headIndex != nil {
		newHeadIndex.Sib = sibMap[headIndex.Next]
		headIndex, newHeadIndex = headIndex.Next, newHeadIndex.Next
	}
	return copyHead
}
