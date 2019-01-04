package main

import "log"

func generateMatrix(n int) [][]int {
	all := n * n
	minI, minJ, maxI, maxJ := 0, 0, n, n
	matrix := make([][]int, 0)
	for i := 0; i < n; i++ {
		matrix = append(matrix, make([]int, n))
	}
	index := 0
	at := "r"
	i, j := 0, 0
	for index < all {
		index++
		matrix[j][i] = index
		switch at {
		case "r":
			if i == maxI-1 {
				minJ++
				at = "d"
				j++
				continue
			}
			i++
		case "d":
			if j == maxJ-1 {
				maxI--
				at = "l"
				i--
				continue
			}
			j++
		case "l":
			if i == minI {
				maxJ--
				at = "u"
				j--
				continue
			}
			i--
		case "u":
			if j == minJ {
				minI++
				at = "r"
				i++
				continue
			}
			j--
		}
	}
	return matrix
}
func main() {
	log.Println(generateMatrix(1))
}