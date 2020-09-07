package _select

import (
	"log"
	"testing"
)

func TestSelectSort(t *testing.T) {
	source1 := []int{0,1,2,3,4,5,6,7,8,9}
	HeapSort(source1)
	log.Println(source1)

}
