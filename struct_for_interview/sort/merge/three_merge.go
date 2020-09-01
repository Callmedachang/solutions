package merge

func ThreeMerge(nums []int)[]int {
	all := make([][]int, len(nums))
	for i, v := range nums {
		all[i] = []int{v}
	}
	for len(all) > 1 {
		if len(all) <= 3 {
			if len(all) == 2 {
				return Merge3(all[0], all[1], []int{})
			}
			if len(all) == 3 {
				return Merge3(all[0], all[1], all[2])
			}
		} else {
			newAll := make([][]int, 0)
			for i := 0; i < len(all); i = i + 3 {
				if i >= len(all)-3 {
					newAll = append(newAll, all[i:]...)
				} else {
					newAll = append(newAll, Merge3(all[i], all[i+1], all[i+2]))
				}
			}
			all = newAll
		}
	}
	return all[0]
}

func Merge3(s1, s2, s3 []int) []int {
	s:=merge(merge(s1, s2), s3)
	return  s
}
