package string

import (
	"log"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	s1 := "1234"
	s2 := "4321"
	log.Println(IsAnagram(s1, s2))
}
