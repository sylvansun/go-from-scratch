package algorithms

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMatrixSum(t *testing.T) {
	Convey("TestMatrixSum", t, func() {

		Convey("single test case", func() {
			nums := [][]int{
				{7, 2, 1},
				{6, 4, 2},
				{6, 5, 3},
				{3, 2, 1},
			}
			expected := 15
			So(MatrixSum(nums), ShouldEqual, expected)
		})
	})
}
