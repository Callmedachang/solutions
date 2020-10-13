package _4spiralOrder

/*
给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。

示例 1:

输入:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
输出: [1,2,3,6,9,8,7,4,5]
示例 2:

输入:
[
  [1, 2, 3, 4],
  [5, 6, 7, 8],
  [9,10,11,12]
]
输出: [1,2,3,4,8,12,11,10,9,5,6,7]
*/
func SpiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	maxTop, maxLeft, maxBottom, maxRight := 0, 0, len(matrix)-1, len(matrix[0])-1
	i, j, index, forward, res, all := 0, 0, 0, 0, make([]int, 0), len(matrix)*len(matrix[0])
	for index < all {
		res = append(res, matrix[i][j])
		switch forward {
		case 0: //从左到右
			if j <= maxRight {
				j++
			}
			if j > maxRight {
				j--
				i++
				forward = 1
				maxTop++
			}
		case 1: //从上到下
			if i <= maxBottom {
				i++
			}
			if i > maxBottom {
				i--
				j--
				forward = 2
				maxRight--
			}
		case 2: //从右到左
			if j >= maxLeft {
				j--
			}
			if j < maxLeft {
				j++
				i--
				forward = 3
				maxBottom--
			}
		case 3: //从下到上
			if i >= maxTop {
				i--
			}
			if i < maxTop {
				i++
				j++
				forward = 0
				maxLeft++
			}
		}
		index++
	}
	return res
}
