package merge

func MergeSort(source []int) []int {
	all := make([][]int, 0)
	for _, v := range source {
		all = append(all, []int{v})
	}
	for len(all) > 1 {
		newAll := make([][]int, 0)
		for i := 0; i < len(all); i = i + 2 {
			if i == len(all)-1 {
				newAll = append(newAll, all[i])
			} else {
				newAll = append(newAll, merge(all[i], all[i+1]))
			}
		}
		all = newAll
	}
	return all[0]
}

func merge(l, r []int) (res []int) {
	lenL, lenR := len(l), len(r)
	for i, j := 0, 0; i < lenL || j < lenR; {
		if i == lenL {
			res = append(res, r[j:]...)
			return
		}
		if j == lenR {
			res = append(res, l[i:]...)
			return
		}
		if l[i] < r[j] {
			res = append(res, l[i])
			i++
		} else {
			res = append(res, r[j])
			j++
		}
	}
	return
}

/*
mark:
将两个有序对数组归并成一个更大的有序数组。通常做法为递归排序，并将两个不同的有序数组归并到第三个数组中。

先来看看动图，归并排序是一种典型的分治应用。
*/
