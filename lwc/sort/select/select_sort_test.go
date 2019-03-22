package _select

import (
	"log"
	"testing"
)

func TestSelectSort(t *testing.T) {
	source1 := []int{9,2,3,4,8,7,6}
	log.Println(SelectSort(source1))
}
