package main

import "fmt"

/*
输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。例如：如果输入如下矩阵：则依次打印出数字1、2、3、4、8、12、16、15、14、13、9、5、6、7、11、10。

*/
func main() {
	s := make([][]int, 5)
	s[0] = []int{1, 2, 3, 4, 5, 6}
	s[1] = []int{1, 2, 3, 4, 5, 6}
	s[2] = []int{1, 2, 3, 4, 5, 6}
	s[3] = []int{1, 2, 3, 4, 5, 6}
	s[4] = []int{1, 2, 3, 4, 5, 6}
	printCycle(s)
}

func printCycle(s [][]int) {
	way, cycleIndex, width, heigh := 1, 0, 0, 0
	hasMore := true
	for hasMore {
		switch way {
		case 1:
			if width > len(s[0]) {
				return
			}
			from, to := 0+cycleIndex, len(s[cycleIndex])-cycleIndex
			for i := from; i < to; i++ {
				fmt.Print(s[cycleIndex][i])
			}
			way++
			width++
		case 2:
			if heigh > len(s) {
				return
			}
			from, to := 1+cycleIndex, len(s)
			for i := from; i < to; i++ {
				fmt.Print(s[i][len(s[0])-cycleIndex-1])
			}
			way++
			heigh++
		case 3:
			if width > len(s[0]) {
				return
			}
			from, to := len(s)-1-cycleIndex, cycleIndex
			for i := from; i >= to; i-- {
				fmt.Print(s[len(s)-cycleIndex-1][i])
			}
			way++
			width++
		case 4:
			if heigh > len(s) {
				return
			}
			from, to := len(s)-1, 1+cycleIndex
			for i := from; i > to; i-- {
				fmt.Print(s[i][cycleIndex])
			}
			way = 1
			cycleIndex++
			heigh++
		}
	}
}
