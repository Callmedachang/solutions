package interview_29_Most_Num_Than_Half

import (
	"log"
	"testing"
)

func TestFindMostNumThanHalf(t *testing.T) {
	s:=[]int{1,2,2,2,2,4}
	log.Println(FindMostNumThanHalf(s))
}