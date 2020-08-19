package main

import (
	"log"
	"sort"
	"strings"
)

/*
给定一个字符串S，检查是否能重新排布其中的字母，使得两相邻的字符不同。

若可行，输出任意可行的结果。若不可行，返回空字符串。

示例 1:

输入: S = "aab"
输出: "aba"
示例 2:

输入: S = "aaab"
输出: ""
注意:
S 只包含小写字母并且长度在[1, 500]区间内。
链接：https://leetcode-cn.com/problems/reorganize-string
*/
func main() {
	log.Println(reorganizeString("aaab"))
}
func reorganizeString(S string) string {
	mm, chars, uniqueChars, newChars,index :=
		make(map[string]int), strings.Split(S, ""), make([]string, 0), make([]string, len(S)),0
	for _, v := range chars {
		if _, ok := mm[v]; !ok {
			uniqueChars = append(uniqueChars, v)
		}
		mm[v]++

	}
	sort.Slice(uniqueChars, func(i, j int) bool {
		return mm[uniqueChars[i]] > mm[uniqueChars[j]]
	})
	if mm[uniqueChars[0]]*2-1> len(S) {
		return ""
	}
	for _, v := range uniqueChars {
		for j := 0; j < mm[v]; j++ {
			if index > len(newChars)-1 {
				index = 1
			}
			newChars[index] = v
			index += 2
		}
	}
	return strings.Join(newChars, "")
}
