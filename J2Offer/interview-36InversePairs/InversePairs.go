package interview_36InversePairs

import "log"
/*
题目：在数组中的两个数字如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。
输入一个数组，求出这个数组中的逆序对的总数。例如在数组{7,5,6,4}中，一共存在5个逆序对，
分别是（7,6）、（7,5）、

*/
func InversePairs(nums []int) {
	MergeSort(nums)
}

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
		if l[i] > r[j] {
			res = append(res, l[i])
			for k := j; k < lenR; k++ {
				log.Println(l[i], r[k])
			}
			i++
		} else {
			res = append(res, r[j])
			j++
		}
	}
	return
}
