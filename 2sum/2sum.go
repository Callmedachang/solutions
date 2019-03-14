package main

import (
	"log"
	"sort"
)

func twosum(sli []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Ints(sli)
	right, left := 0, 0
	max := target - sli[0]
	for i := range sli {
		if sli[len(sli)-1-i] > max {
			i++
		} else {
			right = len(sli) - 1 - i
			break
		}
	}
	for left < right {
		if sli[left]+sli[right] == target {
			res = append(res, []int{sli[left], sli[right]})
			right--
			left++
		} else if sli[left]+sli[right] > target {
			right--
		} else {
			left++
		}
	}
	return res
}
func main() {
	log.Println(twosum([]int{-1, -2, - 3, 0, 1, 2, 3}, 0))
}
