package interview41_sequenceAdd

import "log"

func ContinuesSequenceWithSum(n int) {
	small, big := 1, 2
	for small < (n/2)+1 {
		addRes := add(small, big)
		if addRes == n {
			log.Println(small, big)
			big++
			continue
		}
		if addRes < n {
			big++
		}
		if addRes > n {
			small++
		}
	}
}

func add(from, to int) (res int) {
	for i := from; i <= to; i++ {
		res += i
	}
	return
}
