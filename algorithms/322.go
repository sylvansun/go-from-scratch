package algorithms

import "math"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	lenCoins := len(coins)
	for i := 1; i <= amount; i++ {
		minCount := math.MaxInt32
		for j := 0; j < lenCoins; j++ {
			if i-coins[j] < 0 {
				continue
			} else {
				if dp[i-coins[j]] < minCount {
					minCount = dp[i-coins[j]]
				}
			}
		}
		dp[i] = 1 + minCount
	}
	if dp[amount] > 10005 {
		dp[amount] = -1
	}
	return dp[amount]
}
