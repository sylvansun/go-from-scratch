package algorithms

func KItemsWithMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int {
	ret := 0
	if k <= numOnes {
		ret = k
	} else if k <= numOnes+numZeros {
		ret = numOnes
	} else {
		ret = numOnes - (k - numOnes - numZeros)
	}
	return ret
}
