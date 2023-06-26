package main

import (
	"fmt"
	"math"
)

func pivotInteger(n int) int {
	ret := int(math.Sqrt(float64(n * (n + 1) / 2)))
	if ret*ret == n*(n+1)/2 {
		return ret
	}
	return -1
}

func main() {
	fmt.Println(pivotInteger(8))
}
