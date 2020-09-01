package interview_03_FindInmartix

import (
	"log"
	"testing"
)

func TestFindInMatrix(t *testing.T) {
	s := FindInMatrix([][]int{}, 1)
	log.Println(s)
}
