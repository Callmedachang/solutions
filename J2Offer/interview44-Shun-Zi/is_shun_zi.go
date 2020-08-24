package interview44_Shun_Zi

import "sort"

func JudgeIsShunZi(nums []int) bool {
	sort.Ints(nums)
	zeroNum, noNumCount, before := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroNum++
			before = nums[i]
			continue
		}
		if nums[i]-before > 1 && before != 0 {
			noNumCount += nums[i] - before - 1
		}
		before = nums[i]
	}
	if zeroNum >= noNumCount {
		return true
	}
	return false
}
