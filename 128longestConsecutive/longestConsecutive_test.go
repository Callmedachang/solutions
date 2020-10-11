package _28longestConsecutive

import (
	"log"
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	s:=[]int{100,4,200,1,3,2}
	log.Println(LongestConsecutive(s))
}