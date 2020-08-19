package main

import (
	"log"
	"math/rand"
)

func main() {
	res := make(map[int]int)
	for i := 0; i < 100000; i++ {
		res[rand25()]++
	}

	log.Println(res)
}

func rand25() int {
	return rand5()*5 + rand5()
}

//题目给的条件
func rand5() int {
	return rand.Intn(5)
}
