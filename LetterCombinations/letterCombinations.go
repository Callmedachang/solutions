package main

import (
	"fmt"
	"strings"
)

/*
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。



示例:

输入："23"
输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
说明:
尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。*/

func letterCombinations(digits string) []string {
	numbers := strings.Split(digits, "")
	firstStrs := make([][]string, 0)
	if digits == "" {
		return make([]string, 0)
	}
	for _, number := range numbers {
		switch number {
		case "2":
			firstStrs = append(firstStrs, []string{"a", "b", "c"})
		case "3":
			firstStrs = append(firstStrs, []string{"d", "e", "f"})
		case "4":
			firstStrs = append(firstStrs, []string{"g", "h", "i"})
		case "5":
			firstStrs = append(firstStrs, []string{"j", "k", "l"})
		case "6":
			firstStrs = append(firstStrs, []string{"m", "n", "o"})
		case "7":
			firstStrs = append(firstStrs, []string{"p", "q", "r", "s"})
		case "8":
			firstStrs = append(firstStrs, []string{"t", "u", "v"})
		case "9":
			firstStrs = append(firstStrs, []string{"w", "x", "y", "z"})
		}
	}
	for len(firstStrs) > 1 {
		tempStrs := make([][]string, 0)
		for i := 0; i < len(firstStrs)/2; i++ {
			innerStr := make([]string, 0)
			for j := 0; j < len(firstStrs[i*2]); j++ {
				for k := 0; k < len(firstStrs[i*2+1]); k++ {
					innerStr = append(innerStr, firstStrs[i*2][j]+firstStrs[i*2+1][k])
				}
			}
			tempStrs = append(tempStrs, innerStr)
		}
		if len(firstStrs)%2 == 1 {
			tempStrs = append(tempStrs, firstStrs[len(firstStrs)-1])
		}
		firstStrs = tempStrs
	}
	return firstStrs[0]
}
func main() {
	fmt.Println(letterCombinations("234"))
}
