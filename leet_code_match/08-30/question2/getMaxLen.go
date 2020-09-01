package question2

func GetMaxLen(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	start, ob, max := 0, 1, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 || i == len(nums)-1 {
			if nums[i] == 0 {
				temp := findMax(nums, start, i-1, ob)
				ob=1
				if temp > max {
					max = temp
				}
				start = i + 1
			}
			if i == len(nums)-1 {
				if nums[i] < 0 {
					if ob == 1 {
						ob = -1
					} else {
						ob = 1
					}
				}
				temp := findMax(nums, start, i, ob)
				ob=1
				if temp > max {
					max = temp
				}
				start = i + 1
			}
		} else {
			if nums[i] < 0 {
				if ob == 1 {
					ob = -1
				} else {
					ob = 1
				}
			}
		}
	}
	return max
}
func findMax(nums []int, start, end, ob int) int {
	if end < 0 {
		return 0
	}
	if ob == 1 {
		return end - start + 1
	}
	maxLen := end - start + 1
	needSub := 1
	for end >= start {
		if nums[end] < 0 || nums[start] < 0 {
			return maxLen - needSub
		}
		needSub++
		end--
		start++
	}
	return 0
}
