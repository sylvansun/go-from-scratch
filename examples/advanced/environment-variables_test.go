package advanced

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestEnvironmentVariables(t *testing.T) {

	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println("testing")
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}
