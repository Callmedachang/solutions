package main

import (
	"strings"
)

/*TODO:review
1.典型的滑动窗口结题
2.注意map的复用，而不是每次移动一次游标之后重新循环计算map
*/
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	s1Map, s2Map := make(map[string]int), make(map[string]int)
	s1Slice, s2Slice := strings.Split(s1, ""), strings.Split(s2, "")
	for i := range s1Slice {
		s1Map[s1Slice[i]] += 1
		s2Map[s2Slice[i]] += 1
	}
	if sameEqual(s1Map, s2Map) {
		return true
	}
	for i := len(s1); i < len(s2); i++ {
		s2Map[s2Slice[i]] += 1
		s2Map[s2Slice[i-len(s1)]] -= 1
		if sameEqual(s1Map, s2Map) {
			return true
		}
	}
	return false
}

func sameEqual(s1Map, s2Map map[string]int) bool {
	for k, v := range s2Map {
		if s1Map[k] != v {
			return false
		}
	}
	return true
}

/*
字符串的排列
给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。

换句话说，第一个字符串的排列之一是第二个字符串的子串。

示例1:

输入: s1 = "ab" s2 = "eidbaooo"
输出: True
解释: s2 包含 s1 的排列之一 ("ba").


示例2:

输入: s1= "ab" s2 = "eidboaoo"
输出: False


注意：

输入的字符串只包含小写字母
两个字符串的长度都在 [1, 10,000] 之间*/
