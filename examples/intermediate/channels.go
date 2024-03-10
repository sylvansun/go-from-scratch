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
	go Recv(ch)
	ch <- 10
	fmt.Println("send success")
}

func Recv(c chan int) {
	value := <-c
	fmt.Println("recv success", value)
}

func Producer(factor int, out chan<- int) {
	for i := 0; ; i++ {
		out <- i * factor
	}
}

func Consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
func DoProduceComsume() {
	ch := make(chan int, 64)

	go Producer(1, ch)
	go Producer(2, ch)
	go Consumer(ch)

	time.Sleep(time.Millisecond)
}
