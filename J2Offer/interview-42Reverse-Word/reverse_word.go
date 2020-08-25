package interview_42Reverse_Word

import "strings"

func ReverseWord(s string) string {
	sSli := strings.Split(s, " ")
	for i := 0; i < len(sSli); i++ {
		sSli[i] = Reverse(sSli[i], "")
	}
	return Reverse(strings.Join(sSli, " "), "")
}

func Reverse(s string, split string) string {
	sSli := strings.Split(s, split)
	head, tail := 0, len(sSli)-1
	for head < tail {
		sSli[head], sSli[tail] = sSli[tail], sSli[head]
		head++
		tail--
	}
	return strings.Join(sSli, split)
}
