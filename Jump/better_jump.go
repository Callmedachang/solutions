package main

import "log"

func jump2(nums []int) int {
	i, count, end := 0, 0, len(nums)-1
	var nextI, maxNextI, maxI int
	for i < end {
		if i+nums[i] >= end {
			return count + 1
		}
		nextI, maxNextI = i+1, i+nums[i]
		for nextI <= maxNextI {
			if nextI+nums[nextI] > maxI {
				maxI, i = nextI+nums[nextI], nextI
			}
			nextI++
		}
		count++
	}
	return count
}
func main() {
	log.Println(jump2([]int{3,4,3,2,5,4,3}))
}