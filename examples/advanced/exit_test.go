package advanced

import (
	"os"
	"testing"
)

func TestExitDemo(t *testing.T) {

	defer t.Log("!")

	os.Exit(3)
}
