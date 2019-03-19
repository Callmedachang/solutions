package exchange

func QuickSort(source []int) []int {
	if len(source) <= 1 {
		return source
	}
	target := source[0]
	left, right := make([]int, 0), make([]int, 0)
	for i := 1; i < len(source); i++ {
		cur := source[i]
		if cur < target {
			left = append(left, cur)
			continue
		}
		right = append(right, cur)
	}
	left,right = QuickSort(left),QuickSort(right)
	return append(append(left, target), right...)
}
