package main

import (
	"log"
	"strings"
)

func isMatch(s string, p string) bool {
	//aa
	//*a
	for strings.Contains(p,"**"){
		p=strings.Replace(p,"**","*",-1)
	}
	log.Println(p)
	if p=="*"{
		return true
	}
	sl := strings.Split(s, "")
	pl := strings.Split(p, "")
	return isSliceMatch(sl, pl)

}
func isSliceMatch(s []string, p []string) bool {
	if len(s) == 0 && len(p) == 0 {
		return true
	}
	if len(s) != 0 && len(p) == 0 {
		return false
	}
	if len(s) == 0 && len(p) != 0 {
		for _,v:=range p  {
			if v!="*"{
				return false
			}
		}
		return true
	}
	switch p[0] {
	case "*":
		j := 0
		for isSliceMatch(s[j:], p[1:]) != true {
			j++
			if j > len(s) {
				return false
			}
		}
		return true
	case "?":
		return isSliceMatch(s[1:], p[1:])
	default:
		if s[0] == p[0] {
			return isSliceMatch(s[1:], p[1:])
		} else {
			return false
		}
	}
}
func main() {
	//"abbabaaabbabbaababbabbbbbabbbabbbabaaaaababababbbabababaabbababaabbbbbbaaaabababbbaabbbbaabbbbababababbaabbaababaabbbababababbbbaaabbbbbabaaaabbababbbbaababaabbababbbbbababbbabaaaaaaaabbbbbaabaaababaaaabb"
	//"**aa*****ba*a*bb**aa*ab****a*aaaaaa***a*aaaa**bbabb*b*b**aaaaaaaaa*a********ba*bbb***a*ba*bb*bb**a*b*bb"
	log.Println(isMatch("abbabaaabbabbaababbabbbbbabbbabbbabaaaaababababbbabababaabbababaabbbbbbaaaabababbbaabbbbaabbbbababababbaabbaababaabbbababababbbbaaabbbbbabaaaabbababbbbaababaabbababbbbbababbbabaaaaaaaabbbbbaabaaababaaaabb", "*aa*****ba*a*bb**aa*ab****a*aaaaaa***a*aaaa**bbabb*b*b**aaaaaaaaa*a********ba*bbb***a*ba*bb*bb**a*b*bb"))
}
