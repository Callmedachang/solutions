package main

import (
	"fmt"
	"strings"
)

func generateParenthesis(n int) []string {
	len := 1
	res := make([]string, 0)
	if n==0{
		return res
	}
	res = append(res, "(")
	for len < 2*n {
		temp := make([]string, 0)
		for _, v := range res {
			s1, s2 := nextStr(v,n)
			if s1 != "" {
				temp = append(temp, s1)
			}
			if s2 != "" {
				temp = append(temp, s2)
			}
		}
		res = temp
		len++
	}
	//for i, v := range res {
	//
	//}
	return res
}
func nextStr(sourceStr string, num int) (str1 string, str2 string) {
	arrays := strings.Split(sourceStr, "")
	leftNums := 0
	rightNums := 0
	for _, v := range arrays {
		if v == "(" {
			leftNums++
		} else {
			rightNums++
		}
	}
	if leftNums > rightNums {
		if leftNums < num {
			str1 = sourceStr + "("
			//
			str2 = sourceStr + ")"
		}else{
			str2 = ""
			//
			str1 = sourceStr + ")"
		}

	} else {
		str2 = ""
		//
		str1 = sourceStr + "("
	}
	return
}
func main() {
	fmt.Println(generateParenthesis(1))
}
