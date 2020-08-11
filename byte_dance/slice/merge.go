package main

import (
	"log"
	"sort"
	"math"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	var result [][]int
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result = append(result, intervals[0])
	for i := 1; i < len(intervals); i++ {
		lr := len(result)
		if ret, b := mergeTwo(intervals[i], result[lr-1]); b {
			result[lr-1] = ret
		} else {
			result = append(result, intervals[i])
		}

	}
	return result
}

// 如果可以合并则合并，否则就返回false
func mergeTwo(arr1, arr2 []int) (ret []int, b bool) {
	if arr2[0] > arr1[1] || arr1[0] > arr2[1] {
		return ret, false
	}
	left := int(math.Min(float64(arr1[0]), float64(arr2[0])))
	right := int(math.Max(float64(arr1[1]), float64(arr2[1])))
	return []int{left, right}, true
}

func main() {
	//log.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	//log.Println(merge([][]int{{1, 4}, {0, 5}}))
	//log.Println(merge([][]int{}))
	log.Println(merge([][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}))
}
