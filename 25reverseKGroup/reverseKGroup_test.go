package _5reverseKGroup

import (
	"log"
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	l := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	ll := ReverseKGroup(l, 3)
	log.Println(ll)
}
