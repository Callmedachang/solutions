package _4largestRectangleArea

import (
	"log"
	"testing"
)

func TestFindMaxBefore(t *testing.T) {
	s := []int{4,2,0,3,2,5}
	log.Println(FindMaxBefore(s,5))
}
