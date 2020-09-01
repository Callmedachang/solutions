package match

import "testing"

func TestMatch(t *testing.T) {
	nuts := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	blots := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	Match(nuts,blots)
}
