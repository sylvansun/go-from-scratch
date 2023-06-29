package main

import (
	"fmt"
	"reflect"
)

func basic_func() {
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

func modify_element(array []int, index int) {
	array[index] = 100
	fmt.Println("inner: ", array)
	fmt.Println(reflect.TypeOf(array))
	fmt.Printf("%p\n", &array)    //address of pointers are different
	fmt.Printf("%p\n", &array[0]) //address of elements are the same
}

func test_modify() {
	array := make([]int, 10)
	modify_element(array, 2)
	fmt.Println("outer: ", array)
	fmt.Println(reflect.TypeOf(array))
	fmt.Printf("%p\n", &array)
	fmt.Printf("%p\n", &array[0])
}

func main() {
	// basic_func()
	test_modify()
}
