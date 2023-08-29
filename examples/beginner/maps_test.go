package beginner

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMapNoKey(t *testing.T) {
	Convey("TestMapNoKey", t, func() {

		Convey("single test case\n", func() {
			sampleMap := make(map[uint64]map[uint64][]uint64)
			sampleMap[122] = make(map[uint64][]uint64)
			sampleMap[122][118] = []uint64{1, 2}
			fmt.Println(sampleMap)
			fmt.Println(sampleMap[122])
			fmt.Println(sampleMap[122][118])
			sampleValidList := sampleMap[122][118]
			sampleEmptyList := sampleMap[1][2]
			fmt.Println(len(sampleValidList))
			fmt.Println(len(sampleEmptyList))
		})
	})
}
