package main

import (
	"log"
	"strings"
)

func findSubstring(s string, words []string) []int {
	res := make([]int, 0)
	lenSubStr := 0
	finalStr := ""
	for _, s := range words {
		lenSubStr += len(s)
		finalStr += "#"
	}
	if s == "" || len(words) == 0 {
		return res
	}
	sl := strings.Split(s, "")
	for i := 0; i < len(s)-lenSubStr+1; i++ {
		if isStrEqual(sl[i:i+lenSubStr], words, finalStr) {
			res = append(res, i)
		}
	}
	return res
}
func isStrEqual(sl []string, words []string, fi string) bool {
	s := ""
	for _, w := range sl {
		s += w
	}
	for _, w := range words {
		s = strings.Replace(s, w, "#", 1)
	}
	if len(s) == len(words) && s == fi {
		return true
	} else {
		return false
	}
}

//"wordgoodgoodgoodbestword"
//["word","good","best","good"]
func main() {
	s := "aaa"
	words := []string{"a", "b"}
	log.Println(findSubstring(s, words))
	//
	//s2 := "wordgoodstudentgoodword"
	//words2 := []string{"student", "word"}
	//log.Println(findSubstring(s2, words2))
}
