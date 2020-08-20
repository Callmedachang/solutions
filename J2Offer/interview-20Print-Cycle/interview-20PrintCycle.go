package interview_8Print_Cycle

import "fmt"

func PrintCycle(s [][]int) {
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