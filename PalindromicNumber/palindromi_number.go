package main

import "fmt"

func isPalindrome(x int) bool {
	before:=x
	palindromicNumber := 0
	if x < 0 {
		return false
	}
	for x > 0 {
		palindromicNumber = palindromicNumber*10 + x%10
		x = x / 10
	}
	if before == palindromicNumber {
		return true
	} else {
		return false
	}
}
func main() {
	fmt.Println(isPalindrome(-121))
}