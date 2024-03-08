package main

import (
	"flag"
	"fmt"

	"go-from-scratch/examples/advanced"
)

func main() {
	var hostType, clientMsg, clientMtd string

	flag.StringVar(&hostType, "host", "servr", "host type")
	flag.StringVar(&clientMsg, "msg", "hello", "client message")
	flag.StringVar(&clientMtd, "mtd", "Hello", "client method")
	flag.Parse()

	if hostType == "server" {
		advanced.Serve()
	} else if hostType == "client" {
		advanced.Request(clientMsg)
	} else {
		fmt.Println("Invalid host type")
	}
}
