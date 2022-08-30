package slice

import "testing"

func sliceModify(s []int) {
	s[0] = 100
}

func sliceAppend(s []int) []int {
	s = append(s, 100)
	return s
}


func sliceAppendPtr(s *[]int) {
	*s = append(*s, 100)
	return
}

// 注意：Go 语言中所有的传参都是值传递（传值），都是一个副本，一个拷贝。
// 拷贝的内容是非引用类型（int、string、struct 等这些），在函数中就无法修改原内容数据；
// 拷贝的内容是引用类型（interface、指针、map、slice、chan 等这些），这样就可以修改原内容数据。
func TestSliceFn(t *testing.T) {
	// 参数为引用类型 slice：外层 slice 的 len/cap 不会改变，指向的底层数组会改变
	s := []int{1, 1, 1}
	newS := sliceAppend(s)
	// 函数内发生了扩容
	newNewS := sliceAppend(newS)
	t.Log(s, len(s), cap(s))
	// [1 1 1] 3 3
	t.Log(newS, len(newS), cap(newS))
	// [1 1 1 100] 4 6
	t.Log(newNewS, len(newNewS), cap(newNewS))
	// [1 1 1 100 100] 5 6

	s2 := make([]int, 0, 5)
	newS = sliceAppend(s2)
	// 函数内未发生扩容
	t.Log(s2, s2[0:5], len(s2), cap(s2))
	// [] [100 0 0 0 0] 0 5
	t.Log(newS, newS[0:5], len(newS), cap(newS))
	// [100] [100 0 0 0 0] 1 5

	// 参数为引用类型 slice 的指针：外层 slice 的 len/cap 会改变，指向的底层数组会改变
	sliceAppendPtr(&s)
	t.Log(s, len(s), cap(s))
	// [1 1 1 100] 4 6
	sliceModify(s)
	t.Log(s, len(s), cap(s))
	// [100 1 1 100] 4 6
}
