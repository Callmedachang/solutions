package interview_40FindNumAppearOnce

/*
题目：一个整型数组里除了两个数字之外，其他的数字都出现了两次。请写程序找出这两个只出现一次的数字。
要求时间复杂度是O（n），空间复杂度是O（1）。例如输入数组{2,4,3,6,3,2,5,5}
，因为只有4、6这两个数字只出现一次，其他数字都出现了两次，所以输出4和6。
*/
func FindNumAppearOnce(nums []int) []int {
	start := 0
	for i := 0; i < len(nums); i++ {
		start = start ^ nums[i]
	}
	mask := getDiffOrMask(start)
	head, tail := 0, len(nums)
	for tail > head {
		for nums[tail]&mask != mask {
			tail--
		}
		for nums[head]&mask == mask {
			head++
		}
		nums[head], nums[tail] = nums[tail], nums[head]
	}
	nums1 := nums[0]
	for i := 0; i <= head; i++ {
		nums1 = nums1 ^ nums[i]
	}
	nums2 := nums[head+1]
	for i := head + 1; i < len(nums); i++ {
		nums2 = nums2 ^ nums[i]
	}
	return []int{nums1, nums2}
}

func getDiffOrMask(s int) int {

	return 1
}
