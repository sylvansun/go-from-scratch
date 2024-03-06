package intermediate

import (
	"fmt"
	"time"
)

func Channels() {

	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
}

var cat = make(chan struct{})
var fish = make(chan struct{})
var dog = make(chan struct{})

func Cat() {
	<-cat
	fmt.Println("cat")
	fish <- struct{}{}
}

func Fish() {
	<-fish
	fmt.Println("fish")
	dog <- struct{}{}
}

func Dog() {
	<-dog
	fmt.Println("dog")
	cat <- struct{}{}
}

func doPrintAnimals() {
	for i := 0; i < 100; i++ {
		go Cat()
		go Fish()
		go Dog()
	}
	cat <- struct{}{}
	time.Sleep(2 * time.Second)
}

func NoBufferChannel() {
	ch := make(chan int)
	ch <- 10
	fmt.Println("send success")
}
