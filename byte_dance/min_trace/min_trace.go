package min_trace

func MinTrace(nums [][]int, width, height int) int {
	res := 0
	if width == 1 {
		for _, v := range nums {
			res += v[0]
		}
		return res
	}
	if height == 1 {
		for _, v := range nums[0] {
			res += v
		}
		return res
	}
	return min(MinTrace(nums, width-1, height), MinTrace(nums, width, height-1))
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func MinTrace2(nums [][]int) int {
	res := make([][]int, len(nums))
	for i := range nums {
		res[i] = make([]int, len(nums[0]))
		for j := range nums[i] {
			if i == 0 && j == 0 {
				res[i][j] = nums[i][j]
			}
			if i == 0 {
				res[i][j] = res[i][j-1] + nums[i][j]
			}
			if j == 0 {
				res[i][j] = res[i-1][j] + nums[i-1][j]
			}
			res[i][j] = min(res[i-1][j], nums[i-1][j]) + nums[i][j]
		}
	}
	return res[len(nums)-1][len(nums[0])-1]
}
