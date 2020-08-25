package interview_03_FindInmartix

func FindInMatrix(nums [][]int, target int) bool {
	if len(nums)==0{
		return false
	}
	height, width := len(nums), len(nums[0])
	startI, startJ := width-1, 0
	for startI > 0 || startJ < height {
		if nums[startJ][startI] == target {
			return true
		}
		if nums[startJ][startI] < target {
			if startJ == height-1 {
				return false
			}
			startJ++
		} else {
			if startI == 0 {
				return false
			}
			startI--
		}
	}
	return false
}
