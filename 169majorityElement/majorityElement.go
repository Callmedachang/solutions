package main

import "log"

func majorityElement(nums []int) int {
	all := make(map[int]int)
	max := 0
	maxNum := 0
	for _, b := range nums {
		if all[b] != 0 {
			all[b] = all[b] + 1
		} else {
			all[b] = 1
		}
		if all[b] > maxNum {
			max = b
			maxNum = all[b]
		}
	}
	return max
}
func main() {
	log.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
	log.Println(majorityElement([]int{2, 3, 3}))
}
