package main

import (
	"log"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	res := [][]int{}
	resCopy := [][]int{}
	solution := []int{}
	cs(candidates, solution, target, &res)
	for i := 0; i < len(res); i++ {
		has := false
		for j := i + 1; j < len(res); j++ {
			if isEqual(res[i], res[j]) {
				has = true
				continue
			}
		}
		if !has {
			resCopy = append(resCopy, res[i])
		}
	}
	return resCopy
}
func isEqual(a1, a2 []int) bool {
	res := true
	if len(a1) == len(a2) {
		for k := 0; k < len(a1); k++ {
			if a1[k] != a2[k] {
				res = false
				break
			}
		}
	} else {
		return false
	}
	return res
}
func cs(candidates, solution []int, target int, result *[][]int) {
	if target == 0 {
		*result = append(*result, solution)
	}

	if len(candidates) == 0 || target < candidates[0] {
		// target < candidates[0] 因为candidates是排序好的
		return
	}

	// 这样处理一下的用意是，让切片的容量等于长度，以后append的时候，会分配新的底层数组
	// 避免多处同时对底层数组进行修改，产生错误的答案。
	// 可以注释掉以下语句，运行单元测试，查看错误发生。
	solution = solution[:len(solution):len(solution)]

	cs(candidates[1:], append(solution, candidates[0]), target-candidates[0], result)

	cs(candidates[1:], solution, target, result)
}
func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	s := combinationSum(candidates, 8)
	log.Println(s)

}
