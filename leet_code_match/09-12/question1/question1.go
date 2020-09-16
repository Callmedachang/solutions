package question1

/*
给你一个大小为 rows x cols 的矩阵 mat，其中 mat[i][j] 是 0 或 1，请返回 矩阵 mat 中特殊位置的数目 。

特殊位置 定义：如果 mat[i][j] == 1 并且第 i 行和第 j 列中的所有其他元素均为 0（行和列的下标均 从 0 开始 ），则位置 (i, j) 被称为特殊位置。
*/
func NumSpecial(mat [][]int) int {
	res := 0
	for i := 0; i < len(mat); i++ {
		if index, has := findIndexUnique1(mat[i]); has {
			hasDup := false
			for j := 0; j < len(mat); j++ {
				if mat[j][index] == 0 || j == i {
					continue
				} else {
					hasDup = true
				}
			}
			if !hasDup {
				res++
			}
		}
	}
	return res
}

func findIndexUnique1(nums []int) (int, bool) {
	index := -1
	for i, v := range nums {
		if v == 1 {
			if index >= 0 {
				return 0, false
			} else {
				index = i
			}
		}
	}
	if index < 0 {
		return index, false
	}
	return index, true
}
