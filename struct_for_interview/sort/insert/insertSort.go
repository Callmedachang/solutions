package insert

import (
	"fmt"
)

func InsertSort(a []int) {
	var j int
	for i := 1; i < len(a); i++ {
		temp := a[i]
		for j = i - 1; j >= 0 && a[j] > temp; j-- {
			a[j+1] = a[j]
		}
		a[j+1] = temp
	}
	fmt.Println(a)
}
