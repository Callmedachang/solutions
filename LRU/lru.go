package LRU

import "sync"

type Cache struct {
	size int
	sync.RWMutex
	dict map[string]*Node
	head *Node
	tail *Node
}

type Node struct {
	Key  string
	Val  interface{}
	Prev *Node
	Next *Node
}

func (c *Cache) Set(k string, v interface{}) {
	c.Lock()
	defer c.Unlock()
	if _, has := c.dict[k]; has {
		c.dict[k].Val = v
		c.moveToHead(c.dict[k])
	} else {
		newNode := &Node{Key: k, Val: v}
		c.dict[k] = newNode
		if len(c.dict) > c.size {
			c.removeFromLast()
		}
		c.insertIntoHead(newNode)
	}
}

func (c *Cache) Get(k string) interface{} {
	c.RLock()
	defer c.RUnlock()
	val, has := c.dict[k]
	if has {
		c.moveToHead(c.dict[k])
	}
	return val
}

func (c *Cache) moveToHead(n *Node) {
	if n.Prev == nil {
		return
	}
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
	n.Prev = nil
	n.Next = c.head
	c.head = n
}

func (c *Cache) insertIntoHead(n *Node) {
	if c.head == nil {
		c.head = n
		c.tail = n
	}
	n.Next = c.head
	c.head.Prev = n
	c.head = n
}

func (c *Cache) removeFromLast() {
	if c.tail.Prev == nil {
		c.tail = nil
		c.head = nil
	} else {
		c.tail = c.tail.Prev
		c.tail.Next = nil
	}

}
