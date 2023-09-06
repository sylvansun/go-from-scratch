package library

import (
	"fmt"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

var (
	json = jsoniter.ConfigFastest
)

func TestMapMarShalToString(t *testing.T) {
	sampleMap := make(map[uint64][]uint64)
	sampleMap[1] = []uint64{1, 2, 3}
	sampleMap[99] = []uint64{8, 46, 29}
	str, err := json.MarshalToString(sampleMap)
	fmt.Println(sampleMap)
	fmt.Println(str, err)
	resMap := make(map[uint64][]uint64)
	reserr := json.UnmarshalFromString(str, &resMap)
	fmt.Println(resMap, reserr)
}
