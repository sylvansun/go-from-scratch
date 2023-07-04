package algorithms

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPivotInteger(t *testing.T) {
	Convey("TestPivotInteger", t, func() {

		Convey("single test case", func() {
			number := 8
			expected := 6
			So(PivotInteger(number), ShouldEqual, expected)
		})
	})
}
