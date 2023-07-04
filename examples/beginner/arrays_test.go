package beginner

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTruncatedArray(t *testing.T) {
	Convey("TestTruncatedArray", t, func() {

		Convey("single test case", func() {
			array := make([]int, 5)
			for i := 0; i < 5; i++ {
				array[i] = i
			}
			modified := []int{0, 310, 2, 3, 4}
			expected := []int{310, 2, 3, 4}
			So(TruncatedArray(array), ShouldEqual, expected)
			So(array, ShouldEqual, modified)
		})
	})
}
