package _08sortedArrayToBST

import (
	"log"
	"testing"
)

func TestSortedArrayToBST(t *testing.T) {
	s := []int{-10, -3, 0, 5, 9}
	log.Println(SortedArrayToBST(s))
}
