package mergeSort

import (
	"log"
	"testing"
)

func TestMergeSort(t *testing.T) {
	source1 := []int{1,2,3,4,9,7,6}
	log.Println(MergeSort(source1))
}
