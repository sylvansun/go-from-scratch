package algorithms

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCheckOverlap(t *testing.T) {
	Convey("TestCheckOverlap", t, func() {

		Convey("single test case", func() {
			radius, xcenter, ycenter, x1, y1, x2, y2 := 1, 0, 0, 1, -1, 3, 1
			So(CheckOverlap(radius, xcenter, ycenter, x1, y1, x2, y2), ShouldEqual, true)
		})
	})
}
