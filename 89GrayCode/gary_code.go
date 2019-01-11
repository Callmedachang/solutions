package main

import "log"

/*
0132 2310
0
1
0000
0001
0003
0002

0012
0013
0011
0010

*/

func grayCode(n int) []int {
	if n == 0 {return []int{0}}
	res := []int{0, 1}
	for i := 1; i < n; i++ {
		lenRes := len(res)
		for j := 0; j < lenRes; j++ {
			res = append(res, res[lenRes-1-j]+1<<uint(i))
		}
	}
	return res
}
func main() {
	log.Println(grayCode(3))
}
