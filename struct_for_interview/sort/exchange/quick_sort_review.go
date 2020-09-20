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

func p(nums []int, low, height int) int {
	pivot := nums[height]
	index := low
	for i := low; i < height; i++ {
		if nums[i] < pivot {
			nums[i], nums[index] = nums[index], nums[i]
			index++
		}
	}
	nums[index], nums[height] = nums[height], nums[index]
	return index
}
