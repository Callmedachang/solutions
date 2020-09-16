package main

func NiuDan(n int) int {
	minFlag, maxFlag := 1, 2
	min, max, count := 0, 0, 1
	for {
		min,max = min + minFlag,2*max + maxFlag
		if n >= min && n <= max {
			return count
		}
		count++
		minFlag,maxFlag = minFlag ^ 3,maxFlag ^ 3
	}
}
