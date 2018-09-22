package main

import (
	"strings"
	"fmt"
)

/*
将字符串 "PAYPALISHIRING" 以Z字形排列成给定的行数：

P   A   H   N
A P L S I I G
Y   I   R
之后从左往右，逐行读取字符："PAHNAPLSIIGYIR"

实现一个将字符串进行指定行数变换的函数:

string convert(string s, int numRows);
示例 1:

输入: s = "PAYPALISHIRING", numRows = 3
输出: "PAHNAPLSIIGYIR"
示例 2:

输入: s = "PAYPALISHIRING", numRows = 4
输出: "PINALSIGYAHRPI"
解释:

P     I    N
A   L S  I G
Y A   H R
P     I*/
//that's one function
func convert2(s string, numRows int) string {
	var result string
	if numRows == 1 {
		return s
	}
	sSlice := strings.Split(s, "")
	lenS := len(s)
	tempResult := make([]string, lenS*numRows)
	mid := numRows - 2
	isMid := false
	rankNum := 0
	indexLine := 0
	for len(sSlice) > 0 {
		if isMid == false {
			for i := 0; i < numRows; i++ {
				if len(sSlice) <= 0 {
					break
				}
				tempResult[lenS*i+indexLine] = sSlice[0]
				sSlice = sSlice[1:]
			}
			indexLine++
			isMid = true
		} else {
			for i := 0; i < mid; i++ {
				if len(sSlice) <= 0 {
					break
				}
				rankNum++
				tempResult[lenS*(numRows-i-2)+indexLine] = sSlice[0]
				sSlice = sSlice[1:]
				indexLine++
			}
			rankNum++
			isMid = false
		}
	}
	for _, innerChar := range tempResult {
		if innerChar != "" {
			result += innerChar
		}
	}
	return result
}
func main() {
	fmt.Print(convert2("PAYPALISHIRING", 4))

}

//PIN ALSIG YAHR PI
//PIN ALSIG AYRH PI

/*
P     I    N
A   L S  I G
Y A   H R
P     I*/
