package interview_34Ugly_Number

/*
题目：我们把只包含因子2、3和5的数称作丑数（UglyNumber）。
求按从小到大的顺序的第1500个丑数。例如6、8都是丑数，但14不是，因为
它包含因子7。习惯上我们把1当做第一个丑数。
*/
func GetUglyNumber(n int) int {
	uglyNums := []int{2, 3, 5}
	for len(uglyNums) < n {
		addNums := make([]int, 3)
		for i := 3; i > 1; i-- {
			base := uglyNums[len(uglyNums)-i]
			n2, n3, n5 := base*2, base*3, base*4
			if n2 > uglyNums[len(uglyNums)-1] {
				addNums[3-i] = n2
			} else if n3 > uglyNums[len(uglyNums)-1] {
				addNums[3-i] = n3
			} else {
				addNums[3-i] = n5
			}
		}
		uglyNums = append(uglyNums, addNums...)
	}
	return uglyNums[n-1]
}
