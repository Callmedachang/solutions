package string

import "strings"

/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z 。
*/

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	res := ""
	for i := 0; i < len(strs)-1; i++ {
		if i == 0 {
			res = getCommonPrefix(strs[0], strs[1])
		} else {
			res = getCommonPrefix(strs[i+1], res)
		}
	}
	return res
}

func getCommonPrefix(s1, s2 string) string {
	if s1 == "" || s2 == "" {
		return ""
	}
	res := ""
	s1Slice, s2Slice := strings.Split(s1, ""), strings.Split(s2, "")
	for i := 0; i < len(s1); i++ {
		if i >= len(s2) {
			return res
		}
		if s1Slice[i] == s2Slice[i] {
			res += s1Slice[i]
		} else {
			return res
		}
	}
	return res
}
