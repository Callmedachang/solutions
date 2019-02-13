package main

import (
	"log"
)

type MinStack struct {
	values     []int
	sortValues []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{values: make([]int, 0), sortValues: make([]int, 0)}

}

func (this *MinStack) Push(x int) {
	this.values = append(this.values, x)
	lenSort := len(this.sortValues)
	if lenSort == 0 {
		this.sortValues = append(this.sortValues, x)
		return
	}
	newSorts := make([]int, 0)
	for i := 0; i < lenSort; i++ {
		if x < this.sortValues[i] {
			newSorts=append(newSorts,x)
			newSorts=append(newSorts,this.sortValues[i:]...)
			this.sortValues=newSorts
			return
		}else{
			newSorts=append(newSorts,this.sortValues[i])
		}
	}
	this.sortValues = append(this.sortValues, x)
}

func (this *MinStack) Pop() {
	needPop := this.values[len(this.values)-1]
	this.values = this.values[0 : len(this.values)-1]
	for i := 0; i < len(this.sortValues); i++ {
		if this.sortValues[i] == needPop {
			this.sortValues = append(this.sortValues[0:i], this.sortValues[i+1:]...)
			return
		}
	}
}

func (this *MinStack) Top() int {
	if len(this.values) == 0 {
		return 0
	}
	return this.values[len(this.values)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.values) == 0 {
		return 0
	}
	return this.sortValues[0]
}
func main() {
	s := Constructor()
	s.Push(2)
	s.Push(0)
	s.Push(3)
	s.Push(1)
	log.Println(s.sortValues)
	s.Pop()
	s.Pop()
	s.Pop()
	//log.Println(s.sortValues)
}
