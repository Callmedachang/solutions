package interview_30_Min_K

func MinKWithHeap(nums []int, k int) []int {
	minK := []int{-1}
	for _, v := range nums {
		if len(minK) < k+1 {
			minK = append(minK, v)
			if len(minK) == k+1 {
				HeapSort(minK)
			}
		} else {
			if v < minK[1] {
				minK[1] = v
				//下沉排序
				N := len(minK) - 1
				for N > 1 {
					minK[1], minK[N] = minK[N], minK[1] //将大的放在数组后面，升序排序
					N--
					sink(minK, 1, N)
				}
			}
		}
	}
	return minK
}

func HeapSort(s []int) {
	N := len(s) - 1 //s[0]不用，实际元素数量和最后一个元素的角标都为N
	//构造堆
	//如果给两个已构造好的堆添加一个共同父节点，
	//将新添加的节点作一次下沉将构造一个新堆，
	//由于叶子节点都可看作一个构造好的堆，所以
	//可以从最后一个非叶子节点开始下沉，直至
	//根节点，最后一个非叶子节点是最后一个叶子
	//节点的父节点，角标为N/2
	for k := N / 2; k >= 1; k-- {
		sink(s, k, N)
	}
	//下沉排序
	for N > 1 {
		s[1], s[N] = s[N], s[1] //将大的放在数组后面，升序排序
		N--
		sink(s, 1, N)
	}
}

//下沉（由上至下的堆有序化）
func sink(s []int, k, N int) {
	for {
		i := 2 * k
		if i > N { //保证该节点是非叶子节点
			break
		}
		if i < N && s[i+1] < s[i] { //选择较大的子节点
			i++
		}
		if s[k] <= s[i] { //没下沉到底就构造好堆了
			break
		}
		s[i], s[k] = s[k], s[i]
		k = i
	}
}
