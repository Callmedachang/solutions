package _00_150

import (
	"log"
)

func maxProfit(prices []int) int {
	max := 0
	temp := 0
	for i := 1; i < len(prices); i++ {
		temp += prices[i] - prices[i-1]
		if temp < 0 {
			temp = 0
		}
		if max < temp {
			max = temp
		}
	}
	return max
}
func main() {
	log.Println(maxProfit([]int{1, 4, 1, 4, 3, 1}))
}
