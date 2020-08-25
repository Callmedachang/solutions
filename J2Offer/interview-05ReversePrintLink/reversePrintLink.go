package interview_05ReversePrintLink

import "log"

type Node struct {
	Val  int
	Next *Node
}

func Reverse(head *Node) {
	if head==nil{
		return
	}
	if head.Next == nil {
		log.Println(head.Val)
		return
	}
	Reverse(head.Next)
	log.Println(head.Val)
}
