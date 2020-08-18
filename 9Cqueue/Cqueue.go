package main

import (
	"container/list"
	"log"
)

type CQueue struct {
	l1 *list.List
	l2 *list.List
}

func Constructor() CQueue {
	return CQueue{
		l1: list.New(),
		l2: list.New(),
	}
}

func (c *CQueue) AppendTail(value int) {
	c.l2.PushFront(value)
}

func (c *CQueue) DeleteHead() int {
	if c.l1.Len() == 0 {
		for c.l2.Len() > 0 {
			v := c.l2.Front()
			res := v.Value.(int)
			c.l2.Remove(v)
			c.l1.PushFront(res)
		}
	}
	v := c.l1.Front()
	res := v.Value.(int)
	c.l1.Remove(v)
	return res
}
func main() {
	c:=Constructor()
	c.AppendTail(1)
	c.AppendTail(2)
	c.AppendTail(3)
	log.Println(c.DeleteHead())
	log.Println(c.DeleteHead())
	log.Println(c.DeleteHead())
}