package bucketSort

func sortInBucket(bucket []int) {//此处实现插入排序方式，其实可以用任意其他排序方式
	length := len(bucket)
	if length == 1 {return}

	for i := 1; i < length; i++ {
		backup := bucket[i]
		j := i -1
		//将选出的被排数比较后插入左边有序区
		for  j >= 0 && backup < bucket[j] {//注意j >= 0必须在前边，否则会数组越界
			bucket[j+1] = bucket[j]//移动有序数组
			j -- //反向移动下标
		}
		bucket[j + 1] = backup //插队插入移动后的空位
	}
}
/*
获取数组最大值
 */
func getMaxInArr(arr []int) int{
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max{ max = arr[i]}
	}
	return max
}
/*
桶排序
 */
func sort(arr []int) []int {
	//桶数
	num := len(arr)
	//k（数组最大值）
	max := getMaxInArr(arr)
	//二维切片
	buckets := make([][]int, num)

	//分配入桶
	index := 0
	for i := 0; i < num; i++ {
		index = arr[i] * (num-1) /max//分配桶index = value * (n-1) /k

		buckets[index] = append(buckets[index], arr[i])
	}
	//桶内排序
	tmpPos := 0
	for i := 0; i < num; i++ {
		bucketLen := len(buckets[i])
		if bucketLen > 0{
			sortInBucket(buckets[i])

			copy(arr[tmpPos:], buckets[i])

			tmpPos += bucketLen
		}
	}

	return arr
}