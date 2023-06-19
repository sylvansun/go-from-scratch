package main

import "fmt"

func main() {

	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

	//either adding an additional message or forwarding one will incur deadlock
	limited_msg := make(chan string, 2)
	limited_msg <- "msg1"
	limited_msg <- "msg2"
	// limited_msg <- "msg3"

	fmt.Println(<-limited_msg)
	fmt.Println(<-limited_msg)
	// fmt.Println(<-limited_msg)
}
