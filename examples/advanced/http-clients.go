package main

import (
	"bufio"
	"fmt"
	"net/http"
	"reflect"
)

func main() {

	// resp, err := http.Get("https://gobyexample.com")
	resp, err := http.Get("http://localhost:8090/headers")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println(reflect.TypeOf(resp))
	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
