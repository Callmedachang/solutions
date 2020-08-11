package string

type MinStack struct {
	DataStack []int //数据栈 使用数组实现
	SortStack []int //排序栈 使用数组实现
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{DataStack: make([]int, 0), SortStack: make([]int, 0)}
}

func (m *MinStack) Push(x int) {
	//push的时候数据栈直接放到数组的末尾
	m.DataStack = append(m.DataStack, x)
	//如果排序栈是空的 或者比排序栈的顶栈元素要小于或者等于，那么放到排序栈的栈顶
	if len(m.SortStack) == 0 || (m.SortStack[len(m.SortStack)-1] >= x) {
		m.SortStack = append(m.SortStack, x)
	}
}

func (m *MinStack) Pop() {
	if len(m.DataStack) == 0 {
		return
	}
	//POP的时候数据栈直接弹出顶栈元素
	popItem := m.DataStack[len(m.DataStack)-1]
	m.DataStack = m.DataStack[:len(m.DataStack)-1]
	//如果数据栈的顶栈==排序栈的顶栈元素 那么排序栈的顶栈元素也弹出
	if popItem == m.SortStack[len(m.SortStack)-1] {
		m.SortStack = m.SortStack[:len(m.SortStack)-1]
	}
}

func (m *MinStack) Top() int {
	if len(m.DataStack) == 0 {
		return 0
	}
	return m.DataStack[len(m.DataStack)-1]
}

func (m *MinStack) GetMin() int {
	if len(m.SortStack) == 0 {
		return 0
	}
	return m.SortStack[len(m.SortStack)-1]
}




func main() {
	s := Constructor()
	s.Push(-2)
	s.Push(0)
	s.Push(-3)
	s.GetMin()
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
