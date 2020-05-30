package main

import (
	"log"
	"sync/atomic"
)

var s int32 = 1

func printA() {
	for i := 0; i < 10; {
		if atomic.LoadInt32(&s) == 1 {
			log.Println("A")
			atomic.CompareAndSwapInt32(&s, 1, 2)
			i++
		}
	}

}
func printB() {
	for i := 0; i < 10; {
		if atomic.LoadInt32(&s) == 2 {
			log.Println("B")
			atomic.CompareAndSwapInt32(&s, 2, 3)
			i++
		}
	}
}
func printC() {
	for i := 0; i < 10; {
		if atomic.LoadInt32(&s) == 3 {
			log.Println("C")
			atomic.CompareAndSwapInt32(&s, 3, 1)
			i++
		}
	}
}
func main() {
	go printA()
	go printB()
	go printC()
	ch := make(chan int, 0)
	<-ch
}
