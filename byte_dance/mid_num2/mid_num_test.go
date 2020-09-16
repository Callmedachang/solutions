package mid_num2

import (
	"log"
	"testing"
)

func TestGetMidNum(t *testing.T) {
	s1 := []int{1}
	s2 := []int{2,3,4,5,6}
	log.Println(GetMidNum(s1, s2))
}
