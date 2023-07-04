package algorithms

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMaxSumDivThree(t *testing.T) {
	Convey("TestMaxSumDivThree", t, func() {

		Convey("single test case", func() {
			nums := []int{3, 6, 5, 1, 8}
			expected := 18
			So(MaxSumDivThree(nums), ShouldEqual, expected)
		})
	})
}
