package main

import (
	"fmt"
	"reflect"
)

func basic_func() {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
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
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	modify_element(array, 2)
	fmt.Println("outer: ", array)
	fmt.Println(reflect.TypeOf(array))
	fmt.Printf("%p\n", &array)
	fmt.Printf("%p\n", &array[0])
}

func main() {
	// basic_func()
	test_modify()
	// array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// modify_element(array, 2)
	// fmt.Println(array)
}
