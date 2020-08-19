package exchange_test

import (
	"log"
	"testing"
)

func TestPartition(t *testing.T) {
	ss := []int{7, 4, 6, 3, 7, 2, 9, 9, 3, 4, 6, 7, 45, 5, 6, 3, 8, 7}
	quickSort(ss,0,len(ss)-1)
	log.Println(ss)
}

func qs(s []int) []int {
	if len(s) < 2 {
		return s
	}
	pivot, left, right := s[0], make([]int, 0), make([]int, 0)
	for i := 1; i < len(s); i++ {
		if s[i] <= pivot {
			left = append(left, s[i])
		} else {
			right = append(right, s[i])
		}
	}
	return append(append(qs(left), pivot), qs(right)...)
}

func partition(a []int, lo, hi int) int {
	pivot := a[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if a[j] < pivot {
			i++
			a[j], a[i] = a[i], a[j]
		}
	}
	a[i+1], a[hi] = a[hi], a[i+1]
	return i + 1
}

func quickSort(a []int, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partition(a, lo, hi)
	quickSort(a, lo, p-1)
	quickSort(a, p+1, hi)
}