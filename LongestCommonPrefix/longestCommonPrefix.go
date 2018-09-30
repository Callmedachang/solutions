package main

import (
	"strings"
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	has, length, commonPrefix := getLongestCommonPrefixBetweenTwo(strs[0], strs[1])
	commonPrefixSlice := strings.Split(commonPrefix, "")
	if !has {
		return ""
	}
	for i := 2; i < len(strs); i++ {
		if strs[i]==""{
			return ""
		}
		innerSlice := strings.Split(strs[i], "")
		innerStr := ""
		if len(innerSlice)<length{
			length=len(innerSlice)
		}
		for j := 0; j < length; j++ {
			if commonPrefixSlice[j] == innerSlice[j] {
				innerStr = innerStr + commonPrefixSlice[j]
				continue
			} else {
				if innerStr != "" {
					break
				} else {
					return ""
				}
			}
		}
		commonPrefix = innerStr
		length=len(innerStr)
	}
	return commonPrefix
}
func getLongestCommonPrefixBetweenTwo(str1, str2 string) (has bool, length int, commonPrefix string) {
	str1List := strings.Split(str1, "")
	str2List := strings.Split(str2, "")
	if len(str1) == 0 || len(str2) == 0 {
		return false, 0, ""
	}
	index := 0
	if len(str1) < len(str2) {
		index = len(str1)
	} else {
		index = len(str2)
	}
	for i := 0; i < index; i++ {
		if str1List[i] == str2List[i] {
			commonPrefix = commonPrefix + str1List[i]
			length++
			has = true
		}else{
			break
		}
	}
	return
}
func main() {
	strs:=[]string{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
	"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
	"aab",
	"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
	fmt.Print(longestCommonPrefix(strs))
}