package intermediate

import (
	"fmt"
	"time"
)

func sworker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func ChannelSynchronization() {

	done := make(chan bool, 1)
	go sworker(done)

	fmt.Println(<-done)
	//sending message from channel to string is not allowed
	// var msg string
	// msg <- done
}
