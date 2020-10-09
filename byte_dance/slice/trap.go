package main

/*
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。



上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢 Marcos 贡献此图。

示例:

输入: [0,1,0,2,1,0,1,3,2,1,2,1]
输出: 6
*/
func Trap(height []int) int {
	if len(height)<2{
		return 0
	}
	rightMax, leftMax, rightIndex := make([]int, len(height)), make([]int, len(height)), 0
	leftMax[0] = 0
	res := 0
	for i := 1; i < len(height)-1; i++ {
		if i < rightIndex {
			rightMax[i] = rightMax[i-1]
		} else {
			rightMax[i], rightIndex = findRightMax(height, i)
		}
		if height[i-1] > leftMax[i-1] {
			leftMax[i] = height[i-1]
		} else {
			leftMax[i] = leftMax[i-1]
		}
	}
	for i := range height {
		temp:=min(leftMax[i], rightMax[i]) - height[i]
		if temp>0{
			res +=temp
		}
	}
	return res

}

func findRightMax(nums []int, from int) (val int, index int) {
	for i := from; i < len(nums); i++ {
		if nums[i] > val {
			val = nums[i]
			index = i
		}
	}
	return
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
