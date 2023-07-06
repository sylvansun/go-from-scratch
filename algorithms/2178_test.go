package algorithms

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMaximumEvenSplit(t *testing.T) {
	Convey("TestMaximumEvenSplit", t, func() {

		Convey("single test case", func() {
			finalSum := int64(4)
			expected := []int64{4}
			So(MaximumEvenSplit(finalSum), ShouldEqual, expected)
		})

		Convey("multiple test cases", func() {
			cases := []struct {
				finalSum int64
				expected []int64
			}{
				{8, []int64{2, 6}},
				{20, []int64{2, 4, 6, 8}},
				{30, []int64{2, 4, 6, 8, 10}},
			}
			for _, c := range cases {
				So(MaximumEvenSplit(c.finalSum), ShouldEqual, c.expected)
			}
		})
	})
}
