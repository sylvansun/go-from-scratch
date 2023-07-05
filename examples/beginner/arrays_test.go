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

func TestArrayBoundViolation(t *testing.T) {
	Convey("TestPrint", t, func() {

		emptyArray := ArrayBoundViolation([]int{1, 2, 3, 4, 5})
		emptyExample := make([]int, 0)
		emptyExampleExtra := []int{}
		//these three arrays are the same
		So(emptyArray, ShouldEqual, emptyExample)
		So(emptyArray, ShouldEqual, emptyExampleExtra)
		So(emptyExample, ShouldEqual, emptyExampleExtra)

		So(len(emptyExample), ShouldEqual, 0)
	})
}
