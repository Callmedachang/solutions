package main

import (
	"strconv"
	"strings"
)

/*
10:1
100:9+10^1
1000:19*9 +10^2
*/

func numsOf2(n int) int {
	res := 0
	ns := strings.Split(strconv.Itoa(n), "")
	for i := 0; i < len(ns)-2; i++ {
		in, _ := strconv.Atoi(ns[i])
		if ns[i] != "2" {
			res = res + in*(len(ns)-3) + 10 ^ (len(ns) - 3)
		} else {
			res = res + in*(len(ns)-3) + 1 + getNums(ns[i:])
		}
	}
	return res
}
func getNums(ss []string) int {
	res := 0
	for i := len(ss) - 1; i >= 0; i-- {
		sn, _ := strconv.Atoi(ss[i])
		res = res + sn*(10^(len(ss)-i-1))
	}
	return res
}
