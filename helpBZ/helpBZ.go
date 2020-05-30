package main

import (
	"strings"
	"log"
	"syscall"
	"net"
)

func findFirstCommonChar(str1, str2 string) string {
	//分割2个字符串
	str1Sli,str2Sli := strings.Split(str1, ""),strings.Split(str2, "")
	resMap := make(map[string]*struct{}, 0)
	//第一个字符串数组放入map
	for _, v := range str1Sli {
		resMap[v] = &struct{}{}
	}
	//查找第一个相同的
	for _, v := range str2Sli {
		if resMap[v] != nil {
			return v
		}
	}
	return ""
}
func main() {
	log.Println(findFirstCommonChar("ABCD","ZYJHC"))
	net.Listen()
}