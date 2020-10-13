package _0isValid

import "strings"

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']'的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例2:

输入: "()[]{}"
输出: true
示例3:

输入: "(]"
输出: false
示例4:

输入: "([)]"
输出: false
示例5:

输入: "{[]}"
输出: true
*/
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	slice:=strings.Split(s,"")
	ss := make([]string, 0)
	for _, v := range slice {
		switch v {
		case "(":
			ss = append(ss, v)
		case "[":
			ss = append(ss, v)
		case "{":
			ss = append(ss, v)
		case ")":
			if len(ss) == 0 {
				return false
			}
			if ss[len(ss)-1] == "(" {
				ss = ss[:len(ss)-1]
			} else {
				return false
			}
		case "]":
			if len(ss) == 0 {
				return false
			}
			if ss[len(ss)-1] == "[" {
				ss = ss[:len(ss)-1]
			} else {
				return false
			}
		case "}":
			if len(ss) == 0 {
				return false
			}
			if ss[len(ss)-1] == "{" {
				ss = ss[:len(ss)-1]
			} else {
				return false
			}
		default:
			return false
		}
	}
	if len(ss)==0{
		return true

	}
	return false
}
