package main

import "log"

func spiralOrder(matrix [][]int)[]int {
	res:=make([]int,0)
	if len(matrix) == 0 {
		return res
	}
	minI, minJ, maxI, maxJ := 0, 0, len(matrix[0]), len(matrix)
	all := maxI * maxJ
	index := 0
	at := "r"
	i, j := 0, 0
	for index < all {
		index++
		log.Println(j,i)
		res=append(res,matrix[j][i])
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
	return res
}
func main() {
	//[[1,2,3,4],[5,6,7,8],[9,10,11,12],[13,14,15,16]]
	log.Println(spiralOrder([][]int{{1,2,3,4},{5,6,7,8},{9,10,11,12},{13,14,15,16}}))
}
