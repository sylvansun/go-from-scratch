package algorithms

import "sort"

func MatrixSum(nums [][]int) int {
	ret := 0
	RowNum := len(nums)
	ColumnNum := len(nums[0])
	for i := 0; i < RowNum; i++ {
		sort.Ints(nums[i])
	}
	for j := 0; j < ColumnNum; j++ {
		CurMax := 0
		for i := 0; i < RowNum; i++ {
			if nums[i][j] > CurMax {
				CurMax = nums[i][j]
			}
		}
		ret += CurMax
	}
	return ret
}
