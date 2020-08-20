package interview_30_Min_K

import (
	"log"
	"testing"
)

func TestFindMinK(t *testing.T) {
	s := []int{4, 1, 1, 5, 2, 3}
	ss := MinKWithHeap(s, 4)
	log.Println(ss)
}
