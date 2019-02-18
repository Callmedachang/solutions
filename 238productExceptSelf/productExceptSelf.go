package _38productExceptSelf

/*给定长度为 n 的整数数组 nums，其中 n > 1，返回输出数组 output ，其中 output[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积。

示例:

输入: [1,2,3,4]
输出: [24,12,8,6]
说明: 请不要使用除法，且在 O(n) 时间复杂度内完成此题。

进阶：
你可以在常数空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组不被视为额外空间。）*/
func productExceptSelf(a []int) []int {
	l := len(a)
	left, right := make([]int, l), make([]int, l)

	left[0], right[l-1] = 1, 1
	left[1], right[l-2] = a[0], a[l-1]

	for i := 2; i < l; i++ {
		left[i] = a[i-1] * left[i-1]
		right[l-i-1] = a[l-i] * right[l-i]
	}

	res := make([]int, l)
	for i := 0; i < l; i++ {
		res[i] = left[i] * right[i]
	}
	return res
}
