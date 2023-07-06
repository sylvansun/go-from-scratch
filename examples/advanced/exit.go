package advanced

import (
	"os"
	"testing"
)

func ExitDemo(t *testing.T) {

	defer t.Log("!")

	os.Exit(3)
}
