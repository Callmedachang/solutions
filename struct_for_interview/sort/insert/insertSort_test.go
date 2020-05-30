package insert

import (
	"log"
	"testing"
)

func TestInsertSort(t *testing.T) {
	source := []int{1, 7, 4, 5, 8, 2, 3, 9}
	log.Println(ShellSort(source))
}
