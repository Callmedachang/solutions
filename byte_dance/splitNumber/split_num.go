package splitNumber

import (
	"strings"
)

// splitCount 需要加几次* e.g 6  maxLimit 每个数字最大值 e.g 600
func SplitNum(num string, splitCount int, maxLimit int) int {
	return dp(strings.Split(num, ""), splitCount, maxLimit)
}

func dp(nums []string, splitCount, maxLimit int) int {
	if splitCount == 0 {
		if S2I(nums) > maxLimit {
			return 0
		}
		return 1
	}
	if len(nums) == 0 {
		if splitCount == 0 {
			return 1
		}
		return 0
	}
	res, prefix, index := 0, 0, 1
	for prefix < maxLimit {
		if len(nums) < index+1 {
			return 0
		}
		if res1 := dp(nums[index:], splitCount-1, maxLimit); res1 > 0 {
			res += res1
		}
		index++
		prefix = S2I(nums[:index])
	}
	return res
}

//S2I 字符数组 转  数字  e.g  input = ["1","2","3"]  output = 123
func S2I(nums []string) int {
	res := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if (len(nums)-1)-i == 0 {
			res += i
		} else {
			res += i * 10 * ((len(nums) - 1) - i)
		}
	}
	return res
}
