package intermediate

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func add(num *int) {
	for i := 0; i < 100000; i++ {
		*num += 1
	}
}

func Goroutines() {

	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")
	a := 0
	go add(&a)
	go add(&a)
	fmt.Println(a)
	time.Sleep(time.Second)
	fmt.Println(a)
}
