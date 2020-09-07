package _21maximalSquare

import (
	"log"
	"testing"
)

func TestMaximalSquare(t *testing.T) {
	s := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'}}
	s2 := MaximalSquare(s)
	log.Println(s2)
}
