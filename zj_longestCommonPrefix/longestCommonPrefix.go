package main

import (
	"log"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	beforeSym := ""
	symbol := strings.Split(strs[0], "")
	for j := 0; j < len(symbol); j++ {
		temp := ""
		for k := 0; k <= j; k++ {
			temp = temp + symbol[k]
		}
		for i := 1; i < len(strs); i++ {
			if !strings.HasPrefix(strs[i], temp) {
				return beforeSym
			}
		}
		beforeSym = temp
	}
	return strs[0]
}
func main() {
	log.Println(longestCommonPrefix([]string{"a"}))
}