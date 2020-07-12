package string

import (
	"strings"
	"log"
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

func main() {
	log.Println("S"+reverseWords("  hello world!  "))
}
