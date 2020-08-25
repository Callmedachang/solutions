package interview_08Find_Min_Num

import "log"

/*
题目：把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。
例如数组{3,4,5,1,2}为{1,2,3,4,5}的一个旋转，该数组的最小值为1。
*/
func FindTneMinNum(s []int) int {
	head, tail := 0, len(s)-1
	for {
		mid := (head + tail) / 2
		if s[mid] < s[mid-1] && s[mid] < s[mid+1] {
			return s[mid]
		}
		if s[mid] > s[mid-1] && s[mid] > s[head] {
			head = mid
		}
	}
}
func main() {
	log.Println(FindTneMinNum([]int{3, 5, 6, 8, 9, 1, 2,}))
}
