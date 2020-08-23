package revolve_slice

import (
	"log"
	"testing"
)

func TestRevolveMatrix(t *testing.T) {
	s:=[][]int{{1,2,3},{1,2,3},{1,2,3}}
	log.Println(RevolveMatrix(s))
}