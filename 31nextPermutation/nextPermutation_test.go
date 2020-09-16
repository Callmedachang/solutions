package _1nextPermutation

import (
	"log"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	s := []int{1,3,2}
	NextPermutation(s)
	log.Println(s)
}
