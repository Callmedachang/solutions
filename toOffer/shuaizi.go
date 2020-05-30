package main

import "log"

func P(s, n int) int {
	if s < 1 || n < 1 {
		return 0
	}
	if s > 0 && s < 7 && n == 1 {
		return 1
	}
	return P(s-1, n-1) + P(s-2, n-1) + P(s-3, n-1) + P(s-4, n-1) + P(s-5, n-1) + P(s-6, n-1)
}
func p(s, n int) {
	log.Println(P(s, n), "/", 36)
}
func main() {
	p(6, 2)
}
