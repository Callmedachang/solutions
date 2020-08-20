package interview_30_Min_K

/*
题目：输入n个整数，找出其中最小的k个数。例如输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。
*/
func FindMinK(nums []int, k int) []int {
	return partition(nums, k, []int{})
}

func partition(nums []int, k int, res []int) []int {
	if k == 1 {
		return append(res,findMin(nums))
	}
	pivot := len(nums) / 2
	less, more := make([]int, 0), make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] <= pivot {
			less = append(less, nums[i])
		} else {
			more = append(more, nums[i])
		}
	}
	if len(less) == k {
		return append(res, less...)
	} else if len(less) > k {
		return partition(less, k, []int{})
	} else {
		return partition(more, k-len(less), less)
	}
}


func findMin(nums []int) int {
	cu := nums[0]
	for _, v := range nums {
		if v < cu {
			cu = v
		}
	}
	return cu
}
