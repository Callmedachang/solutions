package main

import (
	"strings"
	"fmt"
)

func longestPalindrome(s string) string {
	lenS := len(s)
	res := ""
	datas := strings.Split(s, "")
	for j := 0; j < lenS; j++ {
		for i := 0; i < lenS-j; i++ {
			if datas[lenS-i-1] == datas[j] {
				_,str:=IsPalindrome(GetSubStr(j, lenS-i-1,datas))
				if len(str)>len(res){
					res= str
				}
			}
		}
	}

	return res
}
func IsPalindrome(sourceStr string) (lenStr int, str string) {
	lenSourceStr := len(sourceStr)
	dataChars := strings.Split(sourceStr, "")
	if lenSourceStr%2 == 1 { //奇数
		for i := 0; i < lenSourceStr/2; i++ {
			if dataChars[i] != dataChars[lenSourceStr-1-i] {
				return 0, ""
			}
		}
	} else { //偶数
		for i := 0; i < lenSourceStr/2; i++ {
			if dataChars[i] != dataChars[lenSourceStr-1-i] {
				return 0, ""
			}
		}
	}
	return lenSourceStr, sourceStr
}
func GetSubStr(start, end int, source []string) (res string) {
	for i := start; i <= end; i++ {
		res += source[i]
	}
	return
}
func main() {
	fmt.Println(longestPalindrome("afffa"))
}
