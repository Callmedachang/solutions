package main

import (
	"log"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < n; i++ {
		nums1[m+i] = nums2[i]
	}
	nums1 = nums1[:m+n]
	sort.Ints(nums1)
}
func main() {
	nums2 := []int{1, 3, 5}
	nums1 := []int{2, 4, 6, 0, 0, 0, 0, 0, 0}
	merge(nums1, 3, nums2, 3)
	log.Println(nums1)
}
