package algorithms

import (
	"math"
)

func PivotInteger(n int) int {
	ret := int(math.Sqrt(float64(n * (n + 1) / 2)))
	if ret*ret == n*(n+1)/2 {
		return ret
	}
	return -1
}
