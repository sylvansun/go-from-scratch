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

func TestAppend(t *testing.T) {
	Convey("Append", t, func() {

		Convey("single test case", func() {
			arrs := [][]int{{1, 2}}
			arr := []int{3, 4, 5}
			fmt.Println(arrs, arr)
			arrs = append(arrs, arr)
			arr[0] = 100
			fmt.Println(arrs, arr)
		})
		Convey("another test case", func() {
			arr := []int{3, 4, 5}
			a := 6
			fmt.Println(arr, a)
			arr = append(arr, a)
			a = 100
			fmt.Println(arr, a)
		})
	})
}
