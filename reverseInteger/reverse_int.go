package main

import "fmt"

const INT_MAX = int(^uint(0) >> 1)

func reverse(x int) int {
	isMinus := false
	index := 10
	result := 0
	if x < 0 {
		isMinus = true
		x = 0 - x
	}
	for x > 0 {
		//溢出判断
		if (result > INT_MAX/10) || (result == INT_MAX/10 && result > 7 && isMinus == false) {
			return 0
		}
		if (result > INT_MAX/10) || (result == INT_MAX/10 && result > 8 && isMinus == true) {
			return 0
		}
		result = result*10 + x%index
		x = x / 10
	}
	if isMinus {
		result = 0 - result
	}
	if result > 2147483647 || result < -2147483648 {
		return 0
	}
	return result
}
func main() {
	fmt.Println(reverse(123))
}
