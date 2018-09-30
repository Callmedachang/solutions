package MaxArea

func MaxArea(height []int) int {
	maxArea := 0
	for i := 0; i < len(height); i++ {
		for j := i; j < len(height); j++ {
			tempMaxArea := getArea(height[i], height[j], i, j)
			if tempMaxArea > maxArea {
				maxArea = tempMaxArea
			}
		}
	}
	return maxArea
}
func getArea(i1, i2, index1, index2 int) int {
	high := 0
	if i1 <= i2 {
		high = i1
	} else {
		high = i2
	}
	result := high * (index2 - index1)
	if result < 0 {
		result = 0 - result
	}
	return result
}
