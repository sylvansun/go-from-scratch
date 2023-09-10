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
			fmt.Println(len(sampleMap[122]))
			sampleMap[122] = make(map[uint64][]uint64)
			fmt.Println(len(sampleMap[122]))
			sampleMap[122][118] = []uint64{1, 2}
			fmt.Println(len(sampleMap[122]))
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

func TestMapTraverse(t *testing.T) {
	Convey("TestMapTraverse", t, func() {

		Convey("single test case\n", func() {
			sampleMap := make(map[uint64][]uint64)
			sampleMap[1] = make([]uint64, 0)
			sampleMap[2] = []uint64{1, 2, 3, 4}

			// for key, value := range sampleMap {
			// 	fmt.Println(key, value)
			// }

			keyRange := 3
			fmt.Println(len(sampleMap))
			for i := 0; i < keyRange; i++ {
				//when visiting a key not set, value will be nil and ok will be false
				value, ok := sampleMap[uint64(i)]
				fmt.Println(i, value, ok)
			}
			boolMap := make(map[uint64]map[uint64]bool)
			boolMap[1] = make(map[uint64]bool)
			boolMap[1][2] = true
			fmt.Println(boolMap[1])
			fmt.Println(boolMap[1][2])
			fmt.Println(boolMap[1][3])
		})
	})
}

func TestEmptyMapVisit(t *testing.T) {
	Convey("TestEmptyMapVisit", t, func() {

		Convey("single test case\n", func() {
			sampleMap := make(map[uint64]map[uint64]map[uint64]bool)
			value1, ok := sampleMap[1]
			fmt.Println(value1, ok) // map[] false
			value2, ok := sampleMap[1][1]
			fmt.Println(value2, ok) // map[] false
			value3, ok := sampleMap[1][1][1]
			fmt.Println(value3, ok) // false false

			// sampleMap[1][1][1] = true // would fail
			// sampleMap[1][1] = map[uint64]bool{1: true} // would fail
			sampleMap[1] = map[uint64]map[uint64]bool{1: {1: true}}
			value1, ok = sampleMap[1]
			fmt.Println(value1, ok) // map[1:map[1:true]] true
			value2, ok = sampleMap[1][1]
			fmt.Println(value2, ok) // map[1:true] true
			value3, ok = sampleMap[1][1][1]
			fmt.Println(value3, ok) // true true
		})
	})
}
