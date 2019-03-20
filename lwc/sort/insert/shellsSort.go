package insert

func ShellSort(a []int) []int {
	n := len(a)
	h := n / 2
	for h >= 1 {
		//将数组变为间隔h个元素有序
		for i := h; i < n; i++ {
			//间隔h插入排序
			for j := i; j >= h; j -= h {
				if a[j] < a[j-h] {
					a[j], a[j-h] = a[j-h], a[j]
				}
			}
		}
		h /= 2
	}
	return a
}
