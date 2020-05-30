package _select

func SelectSort(source []int) []int {
	for i := 0; i < len(source); i++ {
		min, minIndex := source[i], i
		for j := i; j < len(source); j++ {
			if source[j] < min {
				min, minIndex = source[j], j
			}
		}
		source[i], source[minIndex] = source[minIndex], source[i]
	}
	return source
}
