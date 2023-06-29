package main

import (
	"fmt"
	"reflect"
)

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	fmt.Println(iptr)
	fmt.Println(&iptr)
	fmt.Println(reflect.TypeOf(iptr))
	fmt.Println(reflect.TypeOf(&iptr))
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
