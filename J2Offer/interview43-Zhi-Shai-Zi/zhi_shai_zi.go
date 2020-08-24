package interview43_Zhi_Shai_Zi

import "log"

func All(n int) {
	all := Zhi(n)
	resMap := make(map[int]int)
	for _, v := range all {
		resMap[v]++
	}
	allNum:=0
	for _, v := range resMap {
		allNum+=v
	}
	for k, v := range resMap {
		log.Printf("点数【%d】的概率是 【%d/%d】", k, v,allNum)
	}
}

func Zhi(n int) []int {
	if n == 1 {
		return []int{1, 2, 3, 4, 5, 6}
	}
	return add([]int{1, 2, 3, 4, 5, 6}, Zhi(n-1))
}
func add(a, b []int) []int {
	res := make([]int, 0)
	for _, v := range a {
		for _, k := range b {
			res = append(res, v+k)
		}
	}
	return res
}
