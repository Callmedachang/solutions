package interview_39DeepOfTree

import (
	"log"
	"testing"
)

func TestDeepOfTree(t *testing.T) {
	root:=&Node{Val: 1}
	log.Println(DeepOfTree(root))
}