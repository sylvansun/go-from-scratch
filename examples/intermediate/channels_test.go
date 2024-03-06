package intermediate

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDoPrintAnimals(t *testing.T) {
	Convey("TestDoPrintAnimals", t, func() {

		Convey("single test case", func() {
			doPrintAnimals()
		})
	})
}

func TestNoBufferChannel(t *testing.T) {
	Convey("TestNoBufferChannel", t, func() {

		Convey("single test case", func() {
			noBufferChannel()
		})
	})
}
