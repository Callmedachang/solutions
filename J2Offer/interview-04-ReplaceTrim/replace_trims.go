package interview_04_ReplaceTrim

import "strings"

func ReplaceTrims(str string) string {
	newStr := strings.Split(str, "")
	before, trimCount := len(newStr)-1, 0
	for i := range newStr {
		if newStr[i] == " " {
			trimCount++
		}
	}
	newStr = append(newStr, make([]string, trimCount*2)...)
	tail := len(newStr) - 1
	for before >= 0 {
		if newStr[before] != " " {
			newStr[tail] = newStr[before]
			before--
			tail--
		} else {
			newStr[tail], newStr[tail-1], newStr[tail-2] = "3", "2", "%"
			before--
			tail -= 3
		}
	}
	return strings.Join(newStr, "")
}
