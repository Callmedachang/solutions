package exchange

func par(a []int, low, height int) int {
	pivot := a[height]
	index := low
	for i := low; i < height; i++ {
		if a[i] < pivot {
			a[index], a[i] = a[i], a[index]
			index++
		}
	}
	a[index], a[height] = a[height], a[index]
	return index
}
