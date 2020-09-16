package mid_num2

/*
输入两个有序数组，求两个数组中所有数字的中位数，要求时间复杂度为 O(log(m 加 n))，m 和 n 为两个数组的长度。
*/
func GetMidNum(nums1, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	if length%2 == 0 {
		t1 := findK(nums1, nums2, length/2)
		t2 := findK(nums1, nums2, length/2+1)
		return (t1 + t2) / 2
	} else {
		return findK(nums1, nums2, length/2+1)
	}
}

func findK(nums1, nums2 []int, k int) float64 {
	if len(nums1) == 0 {
		return float64(nums2[k-1])
	}
	if len(nums2) == 0 {
		return float64(nums1[k-1])
	}
	if k == 1 {
		if nums1[0] < nums2[0] {
			return float64(nums1[0])
		} else {
			return float64(nums2[0])
		}
	}
	index := k/2
	if len(nums1) < index {
		index = len(nums1)
	}
	if len(nums2) < index {
		index = len(nums2)
	}
	if nums1[index-1] < nums2[index-1] {
		return findK(nums1[index:], nums2, k-index)
	} else {
		return findK(nums1, nums2[index:], k-index)
	}
}
