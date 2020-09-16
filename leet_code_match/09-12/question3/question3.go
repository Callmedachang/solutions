package question3

func MinCostConnectPoints(points [][]int) int {
	if len(points) <= 1 {
		return 0
	}
	res := 0
	for i := 0; i < len(points); i++ {
		index, habs := -1, -1
		for j:=range points[i:] {
			a := findMinAbs(points[j], points)
			if index == -1 {
				index, habs = j, a
			} else {
				if a < habs {
					index, habs = j, a
				}
			}
		}
		points[i], points[index] = points[index], points[i]
		res += habs
	}
	return res
}

func findMinAbs(p []int, points [][]int) int {
	res := -1
	for _, v := range points {
		a := abs(p, v)
		if a==0{
			continue
		}
		if a < res || res == -1 {
			res = a
		}
	}
	return res
}

func abs(n1, n2 []int) int {
	res := 0
	if n1[0]-n2[0] > 0 {
		res += n1[0] - n2[0]
	} else {
		res -= n1[0] - n2[0]
	}

	if n1[1]-n2[1] > 0 {
		res += n1[1] - n2[1]
	} else {
		res -= n1[1] - n2[1]
	}
	return res
}
