package question2

import (
	"log"
	"testing"
)

func TestGetMaxLen(t *testing.T) {
	log.Println(GetMaxLen([]int{1,2,3,5,-6,4,0,10}))
}
