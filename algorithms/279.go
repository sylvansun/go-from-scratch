package algorithms

import "math"

func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		minn := math.MaxInt
		for j := 1; j <= int(math.Sqrt(float64(i))); j++ {
			if dp[i-j*j] < minn {
				minn = dp[i-j*j]
			}
		}
		dp[i] = 1 + minn
	}
	return dp[n]
}
