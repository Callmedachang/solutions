package main

import "fmt"

type IntSlice []int

func (c IntSlice) Len() int {
	return len(c)
}
func (c IntSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c IntSlice) Less(i, j int) bool {
	return c[i] < c[j]
}
func twoSum(nums []int, target int) []int {
	lenNums := len(nums)
	for i := 0; i < lenNums-1; i++ {
		for j := i+1; j < lenNums; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
func main() {
	fmt.Println(twoSum([]int{3,2,4},6))
}