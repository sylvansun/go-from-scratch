package intermediate

import "fmt"

func mayPanic() {
	panic("a problem")
}

func Recover() {

	defer func() {
		if r := recover(); r != nil {

			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()

	fmt.Println("After mayPanic()")
}
