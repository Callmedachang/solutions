package main

import "log"

func solveNQueens(n int)[][]string {
	res := make([][]int, 0)
	init := []int{}
	res = append(res, init)
	for i := 0; i < n; i++ {
		resTemp := make([][]int, 0)
		for _, r := range res {
			tempAll := XX(r, n)
			resTemp = append(resTemp, tempAll...)
		}
		res = resTemp
	}
	resStr := make([][]string, 0)
	for _, v := range res {
		resStr = append(resStr, is2ss(v,n))
	}
	return resStr
}
func XX(s []int, n int) [][]int {
	cur := len(s)
	has := make(map[int]*struct{}, 0)
	res := make([][]int, 0)
	for i := 0; i < cur; i++ {
		distance := cur - i
		if s[i]+distance < n {
			has[s[i]+distance] = &struct{}{}
		}
		if s[i]-distance > -1 {
			has[s[i]-distance] = &struct{}{}
		}
		has[s[i]] = &struct{}{}
	}
	if len(s) == n {
		return append(res, s)
	}
	for i := 0; i < n; i++ {
		if has[i] == nil {
			s1 := append(s, i)
			res = append(res, s1)
		}
	}
	return res
}
func is2ss(s []int,n int) []string {
	res := make([]string, 0)
	for _, v := range s {
		str := ""
		for i := 0; i < n; i++ {
			if i == v {
				str += "Q"
			} else {
				str += "."
			}
		}
		res = append(res,str)
	}
	return res
}
func main() {
	//log.Println(XX([]int{0,3}, 5))
	log.Println(len(solveNQueens(6)))
}
