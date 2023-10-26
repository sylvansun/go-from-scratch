package beginner

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAvgFloat64(t *testing.T) {
	Convey("AvgFloat64", t, func() {

		Convey("single test case", func() {
			sampleSlice := []float64{1.1, 2.2, 3.3, 4.4, 6.5}
			fmt.Println(AvgFloat64(sampleSlice...))
			emptySlice := make([]float64, 0)
			fmt.Println(AvgFloat64(emptySlice...))
		})
	})
}
