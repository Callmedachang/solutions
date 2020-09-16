package mid_num

/*
求数组的中位数。数组由一个升序数组翻转形成，如 1 2 3 4 5 6 7 可以从 5 处翻转，形成 5 6 7 1 2 3 4，求 5 6 7
1 2 3 4 的中位数。要求时间复杂度低于 O(n)。
*/
func GetMidNum(nums []int) int {
	left, right := 0, len(nums)-1
	turnIndex := 0
	for left <= right {
		mid := (left + right) / 2
		if left == right {
			turnIndex = left
			break
		}
		if nums[mid] < nums[mid-1] && nums[mid] < nums[mid+1] {
			turnIndex = mid
			break
		}
		if nums[mid] > nums[0] && nums[mid] > nums[len(nums)-1] {
			if left!=mid{
				left = mid
			}else{
				left ++
			}
			continue
		}
		if nums[mid] > nums[0] && nums[mid] < nums[len(nums)-1] {
			return nums[mid]
		}
		if nums[mid] < nums[0] && nums[mid] < nums[len(nums)-1] {
			right = mid
			continue
		}
	}
	if turnIndex+len(nums)/2>len(nums){
		s:=turnIndex+len(nums)/2-len(nums)
		return nums[s]
	}else{
		s:=turnIndex+len(nums)/2
		return nums[s]
	}
}
