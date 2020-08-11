package main

import "log"

type LCQueue struct {
	size int
	head *node
	tail *node
}
type node struct {
	value int
	next  *node
	prev  *node
}

func Constructor() LCQueue {
	return LCQueue{size: 0}
}

func (this *LCQueue) Push(x int) {
	newNode := &node{value: x}
	if this.size == 0 {
		newNode.prev, newNode.next = nil, nil
		this.head, this.tail = newNode, newNode
	} else if this.size==1{
		this.head.next, this.tail = newNode, newNode
	}else{
		newNode.prev, newNode.next = this.tail, nil
		this.tail = newNode
	}
	this.size++
}

func (this *LCQueue) Pop() {
	if this.size == 0 {
		return
	}
	this.size--
	this.head = this.head.next
}

func (this *LCQueue) Size() int {
	return this.size
}

func (this *LCQueue) Front() int {
	if this.size == 0 {
		return 0
	} else {
		return this.head.value
	}
}

func main() {
	obj := Constructor()
	obj.Size()
	obj.Push(52)
	obj.Front()
	obj.Push(22)
	obj.Front()
	log.Println(obj.Size())
	obj.Pop()
	obj.Front()
	//obj.Pop()
	//obj.Size()
}
