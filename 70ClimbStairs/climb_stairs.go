package main

import "log"

func climbStairs(n int) int {
	b1, b2 := 1, 1
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		for i := 1; i < n; i++ {
			b1, b2 = b2, b1+b2
		}
		return b2
	}
}
func main() {
	log.Println(climbStairs(4))
}
