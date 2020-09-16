package splitNumber

import (
	"log"
	"testing"
)

func TestSplitNum(t *testing.T) {
	s:=SplitNum(" 12322121343434",5,600)
	log.Println(s)
}

func Foo(n int)bool{
	return n&(n-1)==0
}