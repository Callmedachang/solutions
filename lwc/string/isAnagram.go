package string

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	record := make(map[rune]int)
	sl, tl := []rune(s), []rune(t)
	for _, v := range sl {
		record[v] ++
	}
	for _, v := range tl {
		if record[v] <= 0 {
			return false
		}
		record[v] --
	}
	return true
}
