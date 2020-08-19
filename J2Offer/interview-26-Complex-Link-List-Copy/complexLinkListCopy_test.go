package interview_26_Complex_Link_List_Copy

import (
	"log"
	"testing"
)

func TestComplexListNodeClone(t *testing.T) {
	head := &Node{Val: 1}
	head.Next = &Node{Val: 2, Next: &Node{Val: 3}}
	head.Next.Sib = head.Next.Next

	ss := ComplexListNodeClone(head)
	log.Println(ss)
}
