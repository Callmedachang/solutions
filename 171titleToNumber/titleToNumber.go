package main

import (
	"strings"
	"log"
	"math"
)

func main() {
	log.Println(titleToNumber(""))
}
func titleToNumber(s string) int {
	res, sl, w := 0, strings.Split(s, ""), 0
	for i := len(sl) - 1; i >= 0; i-- {
		res = res + int(float64(int([]byte(sl[i])[0] - 64))*math.Pow(26, float64(w)))
		w++
	}
	return res
}
