package _22coinChange

import "math"

/*
给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。

你可以认为每种硬币的数量是无限的。

 

示例 1：

输入：coins = [1, 2, 5], amount = 11
输出：3
解释：11 = 5 + 5 + 1
示例 2：

输入：coins = [2], amount = 3
输出：-1
示例 3：

输入：coins = [1], amount = 0
输出：0
示例 4：

输入：coins = [1], amount = 1
输出：1
示例 5：

输入：coins = [1], amount = 2
输出：2
*/
func coinChange(coins []int, amount int) int {
	return change(0, coins, amount)
}
func change(idx int, coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if idx < len(coins) && amount > 0 {
		max, min := amount/coins[idx], math.MaxInt32
		for i := 0; i <= max; i++ {
			res := change(idx+1, coins, amount-coins[idx]*i)
			if res != -1 {
				min = minVal(res+i, min)
			}
		}
		if min == math.MaxInt32 {
			return -1
		}
		return min
	}
	return -1
}
/*

    private int coinChange(int idxCoin, int[] coins, int amount) {
        if (amount == 0) {
            return 0;
        }
        if (idxCoin < coins.length && amount > 0) {
            int maxVal = amount / coins[idxCoin];
            int minCost = Integer.MAX_VALUE;
            for (int x = 0; x <= maxVal; x++) {
                if (amount >= x * coins[idxCoin]) {
                    int res = coinChange(idxCoin + 1, coins, amount - x * coins[idxCoin]);
                    if (res != -1) {
                        minCost = Math.min(minCost, res + x);
                    }
                }
            }
            return (minCost == Integer.MAX_VALUE)? -1: minCost;
        }
        return -1;
    }
*/
func minVal(a, b int) int {
	if a < b {
		return a
	}
	return b
}
