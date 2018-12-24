package main

import (
	"fmt"
)

func permute(nums []int) [][]int {
	ot := make([][]int, 0)
	if len(nums) == 0 {
		return ot
	}
	if len(nums) == 1 {
		return append(ot, []int{nums[0]})
	}
	ot = append(ot, []int{nums[1]})
	res := make([][]int, 0)
	res = p(nums[0], ot)
	for i := 2; i < len(nums); i++ {
		res = p(nums[i], res)
	}
	return res
}
func p(f int, ot [][]int) [][]int {
	res := make([][]int, 0)
	for _, v := range ot {
		res = append(res, add1(f, v)...)
	}
	return res
}
func add1(r int, is []int) [][]int {
	res := make([][]int, 0)
	for i := 0; i < len(is); i++ {
		res1 := make([]int, 0)
		res1 = append(res1, is[:i]...)
		res1 = append(res1, r)
		res1 = append(res1, is[i:]...)
		res = append(res, res1)
	}
	res = append(res, append(is, r))
	return res
}
func main() {
	l := []int{1, 2, 3}
	fmt.Println(len(permute(l)))
}
