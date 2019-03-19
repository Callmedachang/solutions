package exchange

import (
	"log"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	source := []int{1, 5, 2, 6, 3, 9, 8}
	log.Println(BubbleSort(source))
}
func TestQuickSort(t *testing.T) {
	source := []int{1, 5, 2, 6, 3, 9, 8}
	log.Println(QuickSort(source))
}