package main

import (
	"strings"
	"fmt"
)

//that's one function
func convert(s string, numRows int) string {
	sSlice := strings.Split(s, "")
	tempResult := [9][9]string{}
	mid := numRows - 2
	isMid := false
	rankNum := 0
	for len(sSlice) > 0 {
		if isMid == false {
			for i := 0; i < numRows; i++ {
				if len(sSlice) <= 0 {
					break
				}
				tempResult[i][rankNum] = sSlice[0]
				sSlice = sSlice[1:]
			}
			isMid = true
		} else {
			for i := 0; i < mid; i++ {
				if len(sSlice) <= 0 {
					break
				}
				rankNum++
				tempResult[numRows-i-2][rankNum] = sSlice[0]
				sSlice = sSlice[1:]
			}
			rankNum++
			isMid = false
		}
	}
	fmt.Println(tempResult)
	return "p"
}
func main() {
	convert("PAYPALISHIRING", 4)
}
