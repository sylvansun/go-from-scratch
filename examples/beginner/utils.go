package beginner

func AvgFloat64(input ...float64) float64 {
	sum := float64(0)
	for _, item := range input {
		sum += item
	}
	return sum / float64(len(input))
}
