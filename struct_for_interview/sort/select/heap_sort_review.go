package _select


func HeapSort(s []int) {
	N := len(s) - 1
	for k := N / 2; k >= 1; k-- {
		sink(s, k, N)
	}
	for N > 1 {
		s[1], s[N] = s[N], s[1]
		N--
		sink(s, 1, N)
	}
}

func sink(s []int, k, N int) {
	for {
		i := 2 * k
		if i > N {
			break
		}
		if i < N && s[i+1] > s[i] {
			i++
		}
		if s[k] >= s[i] {
			break
		}
		s[i], s[k] = s[k], s[i]
		k = i
	}
}
