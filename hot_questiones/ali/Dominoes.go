package main

import (
	"fmt"
)

/*
4 Dominoes


Make a chain of dominoes.

Compute a way to order a given set of dominoes in such a way that they form a correct domino chain (the dots on one half of a stone match the dots on the neighbouring half of an adjacent stone) and that dots on the halfs of the stones which don't have a neighbour (the first and last stone) match each other.

For example given the stones 21, 23 and 13 you should compute something like 12 23 31 or 32 21 13 or 13 32 21 etc, where the first and last numbers are the same.

For stones 12, 41 and 23 the resulting chain is not valid: 41 12 23's first and last numbers are not the same. 4 != 3

Some test cases may use duplicate stones in a chain solution, assume that multiple Domino sets are being used.

Input example: (1, 2), (5, 3), (3, 1), (1, 2), (2, 4), (1, 6), (2, 3), (3, 4), (5, 6)
*/
type stones struct {
	Head int
	Tail int
}

func Solution(src []stones) {
	for i := range src {
		tempSrc := copySlice(src)
		tempSrc2 := copySlice(src)
		after := removeIndex(i, tempSrc)
		find([]stones{tempSrc2[i]}, after)
		tempSrc2[i].Change()
		find([]stones{tempSrc2[i]}, after)
	}
}

func find(before, after []stones) {
	head := before[len(before)-1].Tail
	for i, v := range after {
		if v.Head == head || v.Tail == head {
			if v.Tail == head {
				after[i].Change()
			}
			tempBefore, tempAfter := append(before, after[i]), removeIndex(i, after)
			if len(tempAfter) == 0 {
				for _, v := range tempBefore {
					fmt.Print(v)
				}
				continue
			}
			find(tempBefore, tempAfter)
		}
	}
}

func copySlice(src []stones) []stones {
	res := make([]stones, len(src))
	for i := range src {
		res[i] = src[i]
	}
	return res
}

func removeIndex(i int, src []stones) []stones {
	res := src
	res = append(src[0:i], src[i+1:]...)
	return res
}

func (s *stones) Change() {
	s.Tail, s.Head = s.Head, s.Tail
}

func main() {
	//before := []stones{{1, 2}}
	//after := []stones{{2, 3}, {1, 3}}
	//find(before, after)
	Solution([]stones{{2, 1}, {2, 3}, {1, 3}})
}
