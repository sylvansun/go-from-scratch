package algorithms

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNumSquares(t *testing.T) {
	Convey("TestNumSquares", t, func() {

		Convey("single test case added", func() {
			example := 13
			expected := 2
			So(numSquares(example), ShouldEqual, expected)
		})
	})
}

func TestMinFallingPathSum(t *testing.T) {
	Convey("TestMinFallingPathSum", t, func() {

		Convey("single test case", func() {
			matrix := [][]int{
				{2, 1, 3},
				{6, 5, 4},
				{7, 8, 9},
			}
			expected := 13
			So(MinFallingPathSum(matrix), ShouldEqual, expected)
		})
	})
}

func TestReconstructMatrix(t *testing.T) {
	Convey("Given a matrix and upper, lower parameter", t, func() {

		Convey("single test case", func() {
			upper := 5
			lower := 5
			colsum := []int{2, 1, 2, 0, 1, 0, 1, 2, 0, 1}
			expected := [][]int{
				{1, 1, 1, 0, 1, 0, 0, 1, 0, 0},
				{1, 0, 1, 0, 0, 0, 1, 1, 0, 1},
			}
			So(ReconstructMatrix(upper, lower, colsum), ShouldEqual, expected)
		})
	})
}

func TestMaxSumDivThree(t *testing.T) {
	Convey("TestMaxSumDivThree", t, func() {

		Convey("single test case", func() {
			nums := []int{3, 6, 5, 1, 8}
			expected := 18
			So(MaxSumDivThree(nums), ShouldEqual, expected)
		})
	})
}

func TestCheckOverlap(t *testing.T) {
	Convey("TestCheckOverlap", t, func() {

		Convey("single test case", func() {
			radius, xcenter, ycenter, x1, y1, x2, y2 := 1, 0, 0, 1, -1, 3, 1
			So(CheckOverlap(radius, xcenter, ycenter, x1, y1, x2, y2), ShouldEqual, true)
		})
	})
}

func TestMaximumEvenSplit(t *testing.T) {
	Convey("TestMaximumEvenSplit", t, func() {

		Convey("single test case", func() {
			finalSum := int64(4)
			expected := []int64{4}
			So(MaximumEvenSplit(finalSum), ShouldEqual, expected)
		})

		Convey("multiple test cases", func() {
			cases := []struct {
				finalSum int64
				expected []int64
			}{
				{8, []int64{2, 6}},
				{20, []int64{2, 4, 6, 8}},
				{30, []int64{2, 4, 6, 8, 10}},
			}
			for _, c := range cases {
				So(MaximumEvenSplit(c.finalSum), ShouldEqual, c.expected)
			}
		})
	})
}

func TestPivotInteger(t *testing.T) {
	Convey("TestPivotInteger", t, func() {

		Convey("single test case", func() {
			number := 8
			expected := 6
			So(PivotInteger(number), ShouldEqual, expected)
		})
	})
}

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

func TestAlternateDigitSum(t *testing.T) {
	Convey("TestAlternateDigitSum", t, func() {

		Convey("single test case", func() {
			num := 886996
			expected := 0
			So(AlternateDigitSum(num), ShouldEqual, expected)
		})
	})
}

func TestKItemsWithMaximumSum(t *testing.T) {
	Convey("TestKItemsWithMaximumSum", t, func() {

		Convey("single test case", func() {
			numOnes, numZeros, numNegOnes, k := 6, 6, 6, 13
			expected := 5
			So(KItemsWithMaximumSum(numOnes, numZeros, numNegOnes, k), ShouldEqual, expected)
		})
	})
}

func TestMatrixSum(t *testing.T) {
	Convey("TestMatrixSum", t, func() {

		Convey("single test case", func() {
			nums := [][]int{
				{7, 2, 1},
				{6, 4, 2},
				{6, 5, 3},
				{3, 2, 1},
			}
			expected := 15
			So(MatrixSum(nums), ShouldEqual, expected)
		})
	})
}

func TestCoinChange(t *testing.T) {
	Convey("TestCoinChange", t, func() {

		Convey("lc 322", func() {
			coins := []int{1, 2, 5}
			amount := 11
			expected := 3
			So(coinChange(coins, amount), ShouldEqual, expected)
		})
	})
}

func TestNumDistinct(t *testing.T) {
	Convey("TestNumDistinct", t, func() {

		Convey("lc 115", func() {
			s := "rabbbit"
			t := "rabbit"
			expected := 3
			So(numDistinct(s, t), ShouldEqual, expected)
		})
	})
}

func TestCalculateDepth(t *testing.T) {
	Convey("TestCalculateDepth", t, func() {
		Convey("single test case", func() {
			root := &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			}
			expected := 3
			So(calculateDepth(root), ShouldEqual, expected)
		})
	})
}
