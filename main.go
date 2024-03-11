package main

import (
	"flag"
	"fmt"

	"go-from-scratch/examples/advanced"
)

func main() {
	var hostType, clientMsg string
	var CSMode, enableHTTP bool

	flag.StringVar(&hostType, "host", "servr", "host type")
	flag.StringVar(&clientMsg, "msg", "hello", "client message")
	flag.BoolVar(&enableHTTP, "http", false, "server type")
	flag.BoolVar(&CSMode, "cs", false, "client server mode")
	flag.Parse()

	if CSMode {
		if hostType == "server" {
			if enableHTTP {
				advanced.ServeWithHTTP()
			} else {
				advanced.Serve()
			}
		} else if hostType == "client" {
			advanced.Request(clientMsg)
		} else {
			fmt.Println("Invalid host type")
		}
	}
	advanced.DoHandle()
}
