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
	left, right = QuickSort(left), QuickSort(right)
	return append(append(left, target), right...)
}

//quick sort
func QuickSort2(a []int, low, hight int) {
	if low < hight {
		index := Partition(a, low, hight)
		QuickSort2(a, low, index-1)
		QuickSort2(a, index+1, hight)
	}
}

func Partition(a []int, low, hight int) int {
	//指定一个基准元素
	pivot := a[hight]
	//标记交换后的位置
	storeIndex := low
	for i := low; i < hight; i++ {
		if a[i] < pivot {
			a[i], a[storeIndex] = a[storeIndex], a[i]
			storeIndex++
		}
	}
	//最后和基准值交换
	a[hight], a[storeIndex] = a[storeIndex], a[hight]
	return storeIndex
}
//选出第k小元素，k为1~len(s)
func SelectKthMin(s []int, k int) int {
	k--
	lo, hi := 0, len(s)-1
	for {
		j := Partition(s, lo, hi)
		if j < k {
			//对右边继续划分
			lo = j + 1
		} else if j > k {
			//对左边继续划分
			hi = j - 1
		} else {
			return s[k]
		}
	}
}

//选出中位数（比一半的元素小，比另一半的大）
func SelectMid(s []int) int {
	return SelectKthMin(s, len(s)/2)
}