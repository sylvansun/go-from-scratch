package intermediate

import (
	"fmt"
	"time"
)

func DeferPrint() {
	fmt.Println("A")
	defer func() {
		fmt.Println("B")
		defer fmt.Println("C")
		fmt.Println("D")
	}()

	defer fmt.Println("E")
	fmt.Println("F")
}

func DeferLoop() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}

func DeferTime() {
	startedAt := time.Now()
	defer func() {
		fmt.Println(time.Since(startedAt))
	}()

	time.Sleep(time.Second)
}

func DeferTimeClosure() {
	startedAt := time.Now()
	defer func() { fmt.Println(time.Since(startedAt)) }()

	time.Sleep(time.Second)
}
