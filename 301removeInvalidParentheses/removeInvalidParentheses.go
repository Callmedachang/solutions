package main

import (
	"strings"
)

func removeInvalidParentheses(s string) []string {
	res := make([]string, 0)
	abnormalRightParentheses, arIndex, abnormalLeftParentheses, alIndex := 0, -1, 0, -1
	sliceData := strings.Split(s, "")
	for i := 0; i < len(sliceData); i++ {
		if sliceData[i] == "(" {
			abnormalLeftParentheses++
			alIndex = i
		}
		if sliceData[i] == ")" {
			if abnormalLeftParentheses > 0 {
				abnormalLeftParentheses--
			} else {
				abnormalRightParentheses ++
				arIndex = i
			}
		}
	}
	//TODO:在arIndex-alIndex中删除abnormalLeftParentheses个'(' ，在arIndex开始往前删除abnormalRightParentheses个')'
	b := sliceData[:arIndex+1]
	m := sliceData[arIndex+1 : alIndex+1]
	e := sliceData[alIndex+1:]
	sStr := ""
	for _, v := range e {
		sStr = sStr + v
	}
	m1 := del(b, ")", abnormalRightParentheses)
	m2 := del(m, ")", abnormalLeftParentheses)
	add := make(map[string]struct{})
	for k := range m1 {
		for k2 := range m2 {
			add[k+k2+sStr] = struct{}{}
		}
	}
	for k := range add {
		if isValid(k) {
			res = append(res, k)
		}
	}
	return res
}
func isValid(s string) bool {
	abnormalRightParentheses, abnormalLeftParentheses := 0, 0
	sliceData := strings.Split(s, "")
	for i := 0; i < len(sliceData); i++ {
		if sliceData[i] == "(" {
			abnormalLeftParentheses++
		}
		if sliceData[i] == ")" {
			if abnormalLeftParentheses > 0 {
				abnormalLeftParentheses--
			} else {
				abnormalRightParentheses ++
			}
		}
	}
	if abnormalRightParentheses == 0 && abnormalLeftParentheses == 0 {
		return true
	}
	return false
}

func del(sli []string, par string, num int) map[string]struct{} {

	return nil
}

func main() {
	removeInvalidParentheses("))()()()((((()")
}
