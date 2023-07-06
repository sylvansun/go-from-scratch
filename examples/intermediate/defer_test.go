package intermediate

import "testing"

func TestDeferPrint(t *testing.T) {
	DeferPrint()
}

func TestDeferLoop(t *testing.T) {
	DeferLoop()
}

func TestDeferTime(t *testing.T) {
	DeferTime()
}

func TestDeferTimeClosure(t *testing.T) {
	DeferTimeClosure()
}
