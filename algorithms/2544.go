package algorithms

func AlternateDigitSum(n int) int {
	res, sign := 0, 1
	for n > 0 {
		res += n % 10 * sign
		sign = -sign
		n /= 10
	}
	return -sign * res
}
