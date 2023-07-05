package algorithms

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestKItemsWithMaximumSum(t *testing.T) {
	Convey("TestKItemsWithMaximumSum", t, func() {

		Convey("single test case", func() {
			numOnes, numZeros, numNegOnes, k := 6, 6, 6, 13
			expected := 5
			So(KItemsWithMaximumSum(numOnes, numZeros, numNegOnes, k), ShouldEqual, expected)
		})
	})
}
