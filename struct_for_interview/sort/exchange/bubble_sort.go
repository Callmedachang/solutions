package exchange

func BubbleSort(source []int) [] int {
	ok, index := false, 1
	for !ok {
		ok = true
		for i := 0; i < len(source)-index; i++ {
			if source[i] > source[i+1] {
				source[i], source[i+1] = source[i+1], source[i]
				ok = false
			}
		}
		index++
	}
	return source
}

/**
核心：冒泡，持续比较相邻元素，大的挪到后面，因此大的会逐步往后挪，故称之为冒泡。
 */
