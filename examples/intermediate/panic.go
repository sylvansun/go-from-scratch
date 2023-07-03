package intermediate

import (
	"fmt"
	"os"
)

func Panic() {

	// panic("a problem")

	_, err := os.Create("/tmp/file")
	fmt.Println(err)

	if err != nil {
		panic(err)
	}
}
