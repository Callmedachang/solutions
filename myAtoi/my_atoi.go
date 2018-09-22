package main

import (
	"fmt"
	"strings"
)

const INT_MAX = 2147483647
const INT_MIN = 2147483648

func myAtoi(str string) int {
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return 0
	}
	result := 0
	chars := []rune(str)
	isMinus := false
	if chars[0] == '-' {
		isMinus = true
		if len(str) > 1 {
			chars = chars[1:]
		} else {
			return result
		}
	} else if chars[0] == '+' {
		isMinus = false
		if len(str) > 1 {
			chars = chars[1:]
		} else {
			return result
		}
	}
	for _, v := range chars {
		vAscii := int(v)
		if v < 48 || v > 57 {
			if isMinus && result > INT_MIN {
				return 0 - INT_MIN
			}
			if !isMinus && result > INT_MAX {
				return INT_MAX
			}
			if isMinus {
				result = 0 - result
			}
			return result
		} else {
			if isMinus && result > INT_MIN {
				return 0 - INT_MIN
			}
			if !isMinus && result > INT_MAX {
				return INT_MAX
			}
			result = result*10 + vAscii - 48
		}
	}
	if isMinus && result > INT_MIN {
		return 0 - INT_MIN
	}
	if !isMinus && result > INT_MAX {
		return INT_MAX
	}
	if isMinus {
		result = 0 - result
	}
	return result
}
func main() {
	fmt.Println(myAtoi("-42"))
}
