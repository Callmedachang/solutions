package main

import (
	"strings"
)

func kmp(bigStr, smallStr string) bool {
	bs, ss := strings.Split(bigStr, ""), strings.Split(smallStr, "")
	pt, start := getPrefixTable(smallStr), 0
	for i := 0; i < len(bs); {
		if start == len(ss) {
			return true
		}
		if bs[i] == ss[start] {
			start++
			i++
		} else {
			if start = pt[start]; start == -1 {
				i++
				start++
			}
		}
	}
	return false
}

//getPrefixTable:最小前缀匹配表
func getPrefixTable(str string) []int {
	res := make([]int, 0, len(str))
	res = append(res, -1)
	strSli := []rune(str)
	for i := 1; i < len(strSli); i++ {
		res = append(res, getMaxCommonPrefixSuffix(string(strSli[0:i])))
	}
	return res
}

//getMaxCommonPrefixSuffix:最小公共前缀
func getMaxCommonPrefixSuffix(str string) int {
	strSli := []rune(str)
	for i := len(strSli) - 1; i > 0; i-- {
		if strings.HasSuffix(str, string(strSli[0:i])) {
			return i
		}
	}
	return 0
}
