package main

import (
	"fmt"

	"go-from-scratch/examples/advanced"
	"go-from-scratch/examples/intermediate"
)

func main() {
	fmt.Println("hello from main branch")
	fmt.Println("hello from Dev branch")
	fmt.Println(advanced.Add(1, 2))
	intermediate.RateLimiting()
}
