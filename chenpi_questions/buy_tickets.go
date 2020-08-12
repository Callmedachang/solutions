package main

import "log"

//2n个人买票 n 50  n 100 没有预备零钱
// 来一个50的 支持找零次数是
//
//
//
//
//



//c100 代表100的人
//c50 代表50的人
//chipCount 暂存可找零的数量
func buyTickets(c100, c50 int, chipCount int) int {
	if (c100 == 0) || (c50 == 0 && (chipCount == c100)) {
		return 1
	}
	if chipCount == 0 { //没有零钱的前提下，必须要来一个50的 也就是没有想加过程直接往下递归
		if c50 <= 0 {
			return 0
		} else {
			return buyTickets(c100, c50-1, chipCount+1)
		}
	} else { //还有零钱的前提下 可以来100的  也可以来50的所以结果 = 来100的+来50的
		return buyTickets(c100, c50-1, chipCount+1) + buyTickets(c100-1, c50, chipCount-1)
	}
}

func main() {
	log.Println(buyTickets(3, 3, 0))
}
