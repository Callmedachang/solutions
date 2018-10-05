package main

import (
	"fmt"
	"strings"
)

/*给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例 2:

输入: "()[]{}"
输出: true
示例 3:

输入: "(]"
输出: false
示例 4:

输入: "([)]"
输出: false
示例 5:

输入: "{[]}"
输出: true*/
func isValid(s string) bool {
	sArray := strings.Split(s, "")
	leftBracket := make([]string, 0)
	for _, v := range sArray {
		if v == "(" || v == "[" || v == "{" {
			leftBracket = append(leftBracket, v)
		} else if v == ")" || v == "]" || v == "}" {
			if len(leftBracket)<1{
				return false
			}
			lastBracket := leftBracket[len(leftBracket)-1]
			if v == ")" && lastBracket == "(" {
				leftBracket = leftBracket[:len(leftBracket)-1]
				continue
			}
			if v == "]" && lastBracket == "[" {
				leftBracket = leftBracket[:len(leftBracket)-1]
				continue
			}
			if v == "}" && lastBracket == "{" {
				leftBracket = leftBracket[:len(leftBracket)-1]
				continue
			}
			return false
		}
	}
	if len(leftBracket) == 0 {
		return true
	} else {
		return false
	}
}
func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("]"))
}