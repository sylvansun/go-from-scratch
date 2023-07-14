package algorithms

func MinFallingPathSum(matrix [][]int) int {
	ret := 0x3f3f3f3f
	row, col := len(matrix), len(matrix[0])
	for i := 1; i < row; i++ {
		for j := 0; j < col; j++ {
			if j == 0 {
				matrix[i][j] += minInt(matrix[i-1][j], matrix[i-1][j+1])
			} else if j == col-1 {
				matrix[i][j] += minInt(matrix[i-1][j], matrix[i-1][j-1])
			} else {
				matrix[i][j] += minInt(minInt(matrix[i-1][j], matrix[i-1][j-1]), matrix[i-1][j+1])
			}
		}
	}
	for j := 0; j < col; j++ {
		ret = minInt(ret, matrix[row-1][j])
	}
	return ret
}
