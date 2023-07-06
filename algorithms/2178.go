package algorithms

func MaximumEvenSplit(finalSum int64) []int64 {
	ret := []int64{}
	if finalSum%2 > 0 {
		return ret
	}
	for i := int64(2); i <= finalSum; i += 2 {
		ret = append(ret, i)
		finalSum -= i
	}
	ret[len(ret)-1] += finalSum
	return ret
}
