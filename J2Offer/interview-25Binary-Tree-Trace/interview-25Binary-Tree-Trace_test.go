package interview_25Binary_Tree_Trace

import "testing"

func TestRootFirstRange(t *testing.T) {
	s := &Node{Val: 1,
		Left: &Node{Val: 4,
			Left: &Node{Val: 8}, Right: &Node{Val: 9}},
		Right: &Node{Val: 7}}
	BinaryTreeTrace(s,13)
}
