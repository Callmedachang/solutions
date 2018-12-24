package main

import (
	"log"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	all := make(map[string][]string, 0)
	for _, c := range strs {
		if all[sortString(c)] == nil {
			all[sortString(c)] = make([]string, 0)
		}
		all[sortString(c)] = append(all[sortString(c)], c)
	}

	for _, v := range all {
		ri := make([]string, 0)
		for _,v := range v {
			ri = append(ri, v)
		}
		res = append(res, ri)
	}
	return res
}
func sortString(str string) (res string) {
	s := strings.Split(str, "")
	sort.Strings(s)
	for _, v := range s {
		res += v
	}
	return
}
func main() {
	log.Println(groupAnagrams([]string{}))
}
