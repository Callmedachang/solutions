package interview_42_5Reverse_Word2

import "strings"

func ReverseWord2(s string, index int) string {
	if len(s)<index{
		return s
	}
	sSli := strings.Split(s, "")
	sli1, sli2 := sSli[:index], sSli[index:]
	sli1 = Reverse(sli1, "")
	sli2 = Reverse(sli2, "")
	return strings.Join(Reverse(sSli, ""), "")
}

func Reverse(s []string, split string) []string {
	head, tail := 0, len(s)-1
	for head < tail {
		s[head], s[tail] = s[tail], s[head]
		head++
		tail--
	}
	return s
}
