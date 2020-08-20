package interview_8Print_Cycle

import (
	"testing"
)

func TestPrintCycle(t *testing.T) {
	s := make([][]int, 5)
	s[0] = []int{1, 2, 3, 4, 5, 6}
	s[1] = []int{1, 2, 3, 4, 5, 6}
	s[2] = []int{1, 2, 3, 4, 5, 6}
	s[3] = []int{1, 2, 3, 4, 5, 6}
	s[4] = []int{1, 2, 3, 4, 5, 6}
	PrintCycle(s)
}
