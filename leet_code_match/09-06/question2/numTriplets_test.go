package question2

import (
	"log"
	"testing"
)
/*
[4,1,4,1,12]
[3,2,5,4]
*/
func TestNumTriplets(t *testing.T) {
	s := NumTriplets([]int{4,1,4,1,12}, []int{3,2,5,4})
	log.Println(s)
}
