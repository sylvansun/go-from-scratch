package algorithms

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMinFallingPathSum(t *testing.T) {
	Convey("TestMinFallingPathSum", t, func() {

		Convey("single test case", func() {
			matrix := [][]int{
				{2, 1, 3},
				{6, 5, 4},
				{7, 8, 9},
			}
			expected := 13
			So(MinFallingPathSum(matrix), ShouldEqual, expected)
		})
	})
}
