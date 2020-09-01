package match

import "log"

func Match(nuts, blots []int) {
	if len(nuts)==0|| len(blots)==0{
		return
	}
	n := nuts[0]
	b := 0
	smallNuts, bigNuts := make([]int, 0), make([]int, 0)
	smallBlots, bigBlots := make([]int, 0), make([]int, 0)
	for i, v := range blots {
		if blots[i] == n {
			b = blots[i]
			log.Println("nut:", n, ",blot:", v)
			continue
		}
		if blots[i] <= n {
			smallBlots = append(smallBlots, blots[i])
		} else {
			bigBlots = append(bigBlots, blots[i])
		}
	}
	for i := range nuts {
		if nuts[i] == b {
			continue
		}
		if nuts[i] <= b {
			smallNuts = append(smallNuts, nuts[i])
		} else {
			bigNuts = append(bigNuts, nuts[i])
		}
	}
	Match(smallNuts, smallBlots)
	Match(bigNuts, bigBlots)
}
