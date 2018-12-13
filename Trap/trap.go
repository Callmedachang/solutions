package main

import "log"

func trap(height []int) int {
	if len(height) < 2 {
		return 0
	}
	board := make([]int, 0)
	for i := 1; i < len(height)-1; i++ {
		if height[i-1] <= height[i] && height[i] >= height[i+1] {
			board = append(board, i)
		}
	}
	if height[0] > height[1] {
		board = append([]int{0}, board...)
	}
	if height[len(height)-1] > height[len(height)-2] {
		board = append(board, len(height)-1)
	}
	all := 0
	log.Println(board)
	for i := 0; i < len(board)-1; i++ {
		all += getLocation(board[i], board[i+1], i, board, height)
	}
	return all
}
func getLocation(begin, end, index int, board, all []int) int {
	long := end - begin
	trueBegin := getBeginMax(index, board, all)
	trueEnd := getEndMaster(index+1, board, all)
	height := 0
	if trueEnd > trueBegin {
		height = trueBegin
	} else {
		height = trueEnd
	}
	max := long * height
	for i := begin; i < end; i++ {
		if all[i] > height {
			max = max - height
		} else {
			max = max - all[i]
		}
	}
	log.Println(max)
	return max
}
func getBeginMax(begin int, board []int, all []int) int {
	max := all[board[begin]]
	for i := 0; i < begin; i++ {
		if all[board[i]] > max {
			max = all[board[i]]
		}
	}
	return max
}
func getEndMaster(end int, board []int, all []int) int {
	max := all[board[end]]
	for i := len(board)-1; i > end; i-- {
		if all[board[i]] > max {
			max = all[board[i]]
		}
	}
	return max
}
func main() {
	//[5,2,1,2,1,5]
	height := []int{5,3,7,7}
	log.Println(trap(height))
}
