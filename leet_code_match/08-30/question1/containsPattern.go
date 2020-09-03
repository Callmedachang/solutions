package question1

func containsPattern(arr []int, m int, k int) bool {
	if len(arr) < m {
		return false
	}
	for i := 0; i < len(arr); i++ {
		if len(arr)-i < m*k {
			return false
		}
		if k == 1 {
			if m <= len(arr) {
				return true
			}
			return false
		} else {
			for j := 0; j < k-1; j++ {
				start, end := i+j*m, i+j*m+m
				if !sliceEqual(arr[start:end], arr[end:end+m]) {
					break
				} else {
					if j == k-2 {
						return true
					}
				}
			}
		}
	}
	return false
}

func sliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
