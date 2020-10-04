package string

import (
	"log"
	"strings"
)

func reverseWords(s string) string {
	ss := strings.Split(s, " ")
	ss2 := make([]string, 0)
	for i := range ss {
		if strings.TrimSpace(ss[i]) != "" {
			ss2 = append(ss2, strings.TrimSpace(ss[i]))
		}
	}
	for i := 0; i < len(ss2)/2; i++ {
		ss2[i], ss2[len(ss2)-i-1] = ss2[len(ss2)-i-1], ss2[i]
	}

	return strings.Join(ss2, " ")
}

func ReverseWords2(s string) string {
	if len(s) == 0 {
		return s
	}
	for {
		if s[0] == ' ' {
			s = s[1:]
		} else {
			break
		}
	}
	for {
		if s[len(s)-1] == ' ' {
			s = s[:len(s)-1]
		} else {
			break
		}
	}
	s = reverse(s, 0, len(s)-1)
	head, tail := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || i == len(s)-1 {
			if i == len(s)-1 {
				s = reverse(s, head, tail)
			} else {
				s = reverse(s, head, tail-1)
			}

			head, tail = i+1, i+1
		} else {
			tail++
		}
	}
	return s
}

func reverse(s string, begin, end int) string {
	sr := []rune(s)
	for begin < end {
		sr[begin], sr[end] = sr[end], sr[begin]
		begin++
		end--
	}
	return string(sr)
}
func main() {
	log.Println("S" + reverseWords("  hello world!  "))
}
