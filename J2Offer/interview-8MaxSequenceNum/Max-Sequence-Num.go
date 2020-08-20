package interview_8MaxSequenceNum


/*
输入一个整型数组，数组里有正数也有负数。
数组中一个或连续的多个整数组成一个子数组。
求所有子数组的和的最大值。要求时间复杂度为O（n）。
例如输入的数组为{1,2,3,10,4,7,2,5}，和最大的子数组为{3,10,4,7,2}
，因此输出为该子数组的和18。

*/
func MaxSequenceNum(nums []int) int {
	res, current := 0, 0
	for _, v := range nums {
		if current+v > current {
			current += v
			if current > res {
				res = current
			}
		} else {
			if current+v < 0 {
				current = 0
			} else {
				current += v
			}
		}
	}
	return res
}

