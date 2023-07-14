package algorithms

import (
	"math"
)

func MaxSumDivThree(nums []int) int {
	ret := 0
	min11 := math.MaxInt32
	min12 := math.MaxInt32
	min21 := math.MaxInt32
	min22 := math.MaxInt32
	for _, elem := range nums {
		ret += elem
		switch elem % 3 {
		case 1:
			if elem < min11 {
				min12 = min11
				min11 = elem
			} else if elem < min12 {
				min12 = elem
			}
		case 2:
			if elem < min21 {
				min22 = min21
				min21 = elem
			} else if elem < min22 {
				min22 = elem
			}
		}
	}
	switch ret % 3 {
	case 1:
		return ret - minInt(min11, minInt(min12, min21+min22))
	case 2:
		return ret - minInt(min21, min11+min12)
	default:
		return ret
	}
}
