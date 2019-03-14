package _57reverseWords

import "strings"

func reverseWords(s string) string {
	ss := strings.Split(s, " ")
	res := ""
	for _, v := range ss {
		if res == "" {
			res = res + reverseWord(v)
			continue
		}
		res = res + " " + reverseWord(v)
	}
	return res
}

func reverseWord(s string) string {
	sl := strings.Split(s, "")
	sLen := len(sl)
	for i := 0; i < sLen/2; i++ {
		sl[i], sl[sLen-i-1] = sl[sLen-i-1], sl[i]
	}
	res := ""
	for _, v := range sl {
		res = res + v
	}
	return res
}
