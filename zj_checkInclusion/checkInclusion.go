package main

import (
	"log"
	"sort"
	"strings"
)

func checkInclusion(s1 string, s2 string) bool {
	lens1 := len(s1)
	lens2 := len(s2)

	s1Slice := strings.Split(s1, "")
	sort.Strings(s1Slice)
	s1Sorted := ""
	for _, v := range s1Slice {
		s1Sorted = s1Sorted + v
	}

	s2Slice := strings.Split(s2, "")

	for i := 0; i < lens2-lens1+1; i++ {
		temp := subStr(s2Slice, i, lens1)
		if s1Sorted == temp {
			return true
		}
	}
	return false
}

func subStr(str []string, before, long int) string {
	res := make([]string, 0)
	for i := 0; i < long; i++ {
		res = append(res, str[before+i])
	}
	sort.Strings(res)
	resStr := ""
	for _, v := range res {
		resStr = resStr + v
	}
	return resStr
}
func main() {
	log.Println(checkInclusion("hello", "ooolleoooleh"))
}
