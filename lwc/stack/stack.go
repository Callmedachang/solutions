package main

import (
	"log"
	"sync"
)

type LStack struct {
	sync.Locker
	len  int
	data []int
}

func (s *LStack) Pop() int {
	s.Lock()
	defer s.Unlock()
	if s.len == 0 {
		return 0
	}
	res := s.data[s.len-1]
	s.data = s.data[0 : s.len-1]
	s.len--
	return res
}

func (s *LStack) Push(v int) {
	s.Lock()
	defer s.Unlock()
	s.data = append(s.data, v)
	s.len++
}

func newLStack() *LStack {
	return &LStack{len: 0, data: make([]int, 0)}
}

func main() {
	s := newLStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	log.Println(s.Pop())
	log.Println(s.Pop())
	log.Println(s.Pop())
}
