package interview_27_BST_To_DoubleLinkList

import (
	"log"
	"testing"
)

func TestBSTToDoubleLinkList(t *testing.T) {
	root := &Node{Val: 5, Left: &Node{Val: 3, Left: &Node{Val: 2}, Right: &Node{Val: 4}}, Right: &Node{Val: 6}}
	ss := BSTToDoubleLinkList(root)
	log.Println(ss)

}
