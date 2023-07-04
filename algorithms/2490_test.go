package algorithms

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIsCircularSentence(t *testing.T) {
	Convey("TestIsCircularSentence", t, func() {
		Convey("single test case", func() {
			demo := "A man, a plan, a canal: Panama"
			So(IsCircularSentence(demo), ShouldBeFalse)
		})
		//Sub Conveys are all executed, even if one of them fails.
		//But So assertions will stop if the previous one fails.
		Convey("multiple test cases", func() {
			cases := []struct {
				Sentence string
				Expected bool
			}{
				{"bad dog", false},
				{"simple egg got twist toss", true},
				{"asdfiahds ifqhf", false},
			}
			for _, c := range cases {
				So(IsCircularSentence(c.Sentence), ShouldEqual, c.Expected)
			}
		})

	})

}
