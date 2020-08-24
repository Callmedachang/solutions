package interview42Reverse_Word

import (
	"log"
	"testing"
)

func TestReverseWord(t *testing.T) {
	log.Println(Reverse("i am a student"," "))
	log.Println(ReverseWord("i am a student"))
}