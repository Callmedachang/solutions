package _4longestCommonPrefix

/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:
*/

func LongestCommonPrefix(strs []string) string {
	maxPre := ""
	for i, v := range strs {
		if i == 0 {
			maxPre = v
		} else {
			pre := FindCommonPre(maxPre, v)
			if pre == "" {
				return ""
			}
			if len(pre) < len(maxPre) {
				maxPre = pre
			}
		}
	}
	return maxPre
}

func FindCommonPre(s1, s2 string) string {
	if len(s1) < len(s2) {
		for i := len(s1); i > 0; i-- {
			if s1[:i] == s2[:i] {
				return s1[:i]
			}
		}
	} else {
		for i := len(s2); i > 0; i-- {
			if s2[:i] == s1[:i] {
				return s2[:i]
			}
		}
	}
	return ""
}
