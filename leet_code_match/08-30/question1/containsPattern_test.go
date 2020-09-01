package question1

import (
	"log"
	"testing"
)

func TestContainsPattern(t *testing.T) {
	s:=[]int{1}
	ss:=ContainsPattern(s,8,3)
	log.Println(ss)
}