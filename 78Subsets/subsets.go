package main

import "log"

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	for _, v := range nums {
		temp := make([][]int, 0)
		for i := 0; i < len(res); i++ {
			temp = append(temp, append(res[i], v))
		}
		res = append(res, temp...)
		res = append(res, []int{v})
	}
	res = append(res, []int{})
	return res
}
func main() {
	log.Println(subsets([]int{1}))
}
