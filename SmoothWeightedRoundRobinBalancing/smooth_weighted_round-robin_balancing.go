package main

import (
	"log"
	"time"
)

type node struct {
	Name         string
	CurrentValue int
	Value        int
}

func SWRBBalance(nodes []*node) {
	totalBalance := getTotalBalance(nodes)
	i := 1
	for {
		//like netWork I/O
		time.Sleep(time.Second)
		s := selectNode(nodes, totalBalance)
		//print selectNode
		log.Println("no.", i, "--selected nodeName:", s.Name)
		//index
		i++
	}
}

func getTotalBalance(nodes []*node) (res int) {
	for _, v := range nodes {
		res = res + v.Value
	}
	return
}

func selectNode(nodes []*node, totalBalance int) (res *node) {
	nodeIndex := 0
	for i := 0; i < len(nodes); i++ {
		if res == nil {
			res = nodes[i]
		} else {
			if nodes[i].CurrentValue > res.CurrentValue {
				res = nodes[i]
				nodeIndex = i
			}
		}
	}
	for i := 0; i < len(nodes); i++ {
		nodes[i].CurrentValue += nodes[i].Value
	}
	nodes[nodeIndex].CurrentValue -= totalBalance
	return res
}

func main() {
	n1 := &node{"A-1", 1, 1}
	n2 := &node{"B-2", 2, 2}
	n3 := &node{"C-3", 3, 3}
	nodes := make([]*node, 0)
	nodes = append(nodes, n1, n2, n3)
	SWRBBalance(nodes)
}
