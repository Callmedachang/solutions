package revolve_slice

func RevolveMatrix(nums [][]int) {
	l, index := len(nums)-1, len(nums)/2+len(nums)%1
	for i := 0; i <= index; i++ {
		innerIndex := (len(nums)-2*i)/2 + (len(nums)-2*i)%1
		for j := i; j <= innerIndex; j++ {
			temp1, temp2, temp3 := 0, 0, 0
			temp1, nums[i][j] = nums[i][j], nums[j][l-i]
			temp2, nums[l-j][i] = nums[l-j][i], temp1
			temp3, nums[l-i][l-j] = nums[l-i][l-j], temp2
			nums[j][l-i] = temp3
		}
	}
}