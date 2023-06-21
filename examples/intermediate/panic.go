package main

import (
	"fmt"
	"os"
)

func main() {

	// panic("a problem")

	_, err := os.Create("/tmp/file")
	fmt.Println(err)

	if err != nil {
		panic(err)
	}
}
