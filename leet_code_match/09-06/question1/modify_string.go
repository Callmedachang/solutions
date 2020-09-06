package question1

func ModifyString(s string) string {
	if s == "?" {
		return "a"
	}
	sSli := []byte(s)
	for i, v := range sSli {
		if v == '?' {
			if i == 0 {
				var temp byte = 'a'
				for {
					if temp != sSli[i+1] {
						sSli[i] = temp
						break
					} else {
						temp++
					}
				}
			}
			if i > 0 && i < len(s)-1 {
				var temp byte = 'a'
				for {
					if temp != sSli[i+1] && temp != sSli[i-1] {
						sSli[i] = temp
						break
					} else {
						temp++
					}
				}
			}
			if i == len(s)-1 {
				var temp byte = 'a'
				for {
					if temp != sSli[i-1] {
						sSli[i] = temp
						break
					} else {
						temp++
					}
				}
			}
		}
	}
	return string(sSli)
}
