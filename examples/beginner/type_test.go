package beginner

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestString(t *testing.T) {
	Convey("TestString", t, func() {

		Convey("single test case", func() {
			name := "Eve"
			expected := 5
			So(Eve.String(), ShouldEqual, name)
			So(Eve, ShouldEqual, expected)
		})
	})
}
