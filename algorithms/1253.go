package algorithms

func ReconstructMatrix(upper int, lower int, colsum []int) [][]int {
	ret := make([][]int, 2)
	for i := 0; i < 2; i++ {
		ret[i] = make([]int, len(colsum))
	}
	unit_indices := make([]int, 0, len(colsum))
	for index, value := range colsum {
		switch value {
		case 2:
			ret[0][index], ret[1][index] = 1, 1
			upper -= 1
			lower -= 1
			if upper < 0 || lower < 0 {
				return make([][]int, 0)
			}
		case 1:
			unit_indices = append(unit_indices, index)
		}
	}
	if upper+lower != len(unit_indices) {
		return make([][]int, 0)
	} else {
		for i := 0; i < len(unit_indices); i++ {
			if upper > 0 {
				ret[0][unit_indices[i]] = 1
				upper -= 1
			} else {
				ret[1][unit_indices[i]] = 1
			}
		}
	}
	return ret
}
