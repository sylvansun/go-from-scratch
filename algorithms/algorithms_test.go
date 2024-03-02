package algorithms

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNumSquares(t *testing.T) {
	Convey("TestNumSquares", t, func() {

		Convey("single test case", func() {
			example := 13
			expected := 2
			So(numSquares(example), ShouldEqual, expected)
		})
	})
}
