package algorithms

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReconstructMatrix(t *testing.T) {
	Convey("Given a matrix and upper, lower parameter", t, func() {

		Convey("single test case", func() {
			upper := 5
			lower := 5
			colsum := []int{2, 1, 2, 0, 1, 0, 1, 2, 0, 1}
			expected := [][]int{
				{1, 1, 1, 0, 1, 0, 0, 1, 0, 0},
				{1, 0, 1, 0, 0, 0, 1, 1, 0, 1},
			}
			So(ReconstructMatrix(upper, lower, colsum), ShouldEqual, expected)
		})
	})
}
