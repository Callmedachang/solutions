package merge

import (
	"log"
	"testing"
)

func TestMergeSort(t *testing.T) {
	source1 := []int{1,2,3,4,9,7,6}
	log.Println(ThreeMerge(source1))
}

func TestMerge3(t *testing.T) {
	log.Println(Merge3([]int{5},[]int{3},[]int{4}))
}