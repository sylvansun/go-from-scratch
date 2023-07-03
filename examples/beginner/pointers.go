package beginner

import (
	"fmt"
	"reflect"
)

func zeroval(ival int) {
	fmt.Println(ival)
	ival = 0
}

func zeroptr(iptr *int) {
	fmt.Println(iptr)
	fmt.Println(&iptr)
	fmt.Println(reflect.TypeOf(iptr))
	fmt.Println(reflect.TypeOf(&iptr))
	*iptr = 0
}

func Pointers() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
