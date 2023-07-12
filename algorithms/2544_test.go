package algorithms

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAlternateDigitSum(t *testing.T) {
	Convey("TestAlternateDigitSum", t, func() {

		Convey("single test case", func() {
			num := 886996
			expected := 0
			So(AlternateDigitSum(num), ShouldEqual, expected)
		})
	})
}
