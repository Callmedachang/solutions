package interview_29_Most_Num_Than_Half

/*

题目：数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。例如输入一个长度为9的数组{1,2,3,2,2,2,5,4,2}。
由于数字2在数组中出现了5次，超过数组长度的一半，因此输出2。
*/
func FindMostNumThanHalf(nums []int) int {
	currentNum, count := 0, 0
	for _, v := range nums {
		if v == currentNum {
			count++
			continue
		} else {
			if count > 1 {
				count--
			} else {
				count,currentNum = 1,v
			}
		}
	}
	return currentNum
}
