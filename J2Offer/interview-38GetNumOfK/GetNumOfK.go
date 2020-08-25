package interview_38GetNumOfK

/*
数字在排序数组中出现的次数题目：统计一个数字在排序数组中出现的次数。
例如输入排序数组{1,2,3,3,3,3,4,5}和数字3，由于3在这个数组中出现了4次，因此输出4。
*/
func GetNumOfK(nums []int, k int) int {
	start := GetIndexMin(nums, 0, len(nums)-1, k)
	end := GetIndexMax(nums, 0, len(nums)-1, k)
	if start == end && start == -1 {
		return -1
	}
	if start == end {
		return start
	}
	return end - start + 1
}

func GetIndexMin(nums []int, start, end, target int) int {
	if end-start <= 1 {
		if nums[start] == target {
			return start
		}
		if nums[end] == target {
			return end
		}
		return -1
	}
	mid := (start + end) / 2
	if nums[mid] == target {
		if mid == 0 {
			return 0
		}
		if nums[mid-1] < target {
			return mid
		}
		return GetIndexMin(nums, start, mid, target)
	}
	if nums[mid] > target {
		return GetIndexMin(nums, start, mid, target)
	} else {
		return GetIndexMin(nums, mid, end, target)
	}
}

func GetIndexMax(nums []int, start, end, target int) int {
	if end-start <= 1 {
		if nums[end] == target {
			return end
		}
		if nums[start] == target {
			return start
		}
		return -1
	}
	mid := (start + end) / 2
	if nums[mid] == target {
		if mid == end {
			return mid
		}
		if nums[mid+1] > target {
			return mid
		}
		return GetIndexMax(nums, mid, end, target)
	}
	if nums[mid] > target {
		return GetIndexMax(nums, start, mid, target)
	} else {
		return GetIndexMax(nums, mid, end, target)
	}
}
