package _select

func HeapSort(s []int) {
	N := len(s) - 1
	for k := N / 2; k >= 0; k-- {
		sink(s, k, N)
	}
	for N > 0 {
		s[0], s[N] = s[N], s[0]
		N--
		sink(s, 0, N)
	}
}

func sink(s []int, k, N int) {
	for {
		i := 2*k + 1
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
