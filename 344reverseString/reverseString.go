package _44reverseString

func reverseString(s []byte) {
	sLen := len(s)
	for i := 0; i < sLen/2; i++ {
		s[i], s[sLen-i-1] = s[sLen-i-1], s[i]
	}
	return
}
