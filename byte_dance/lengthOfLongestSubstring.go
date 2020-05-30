package main

import (
	"strings"
	"log"
)

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	strSlice, tempMap, max := strings.Split(s, ""), make(map[string]int), 0
	for i, v := range strSlice {
		if _, ok := tempMap[v]; ok {
			cutMap(tempMap, tempMap[v])
		}
		tempMap[v] = i
		if len(tempMap) > max {
			max = len(tempMap)
		}
	}
	return max
}

func cutMap(src map[string]int, index int) {
	for k, v := range src {
		if v <= index {
			delete(src, k)
		}
	}
	return
}

func main() {
	log.Println(lengthOfLongestSubstring("abccc"))
}
