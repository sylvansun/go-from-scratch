package beginner

import (
	"fmt"
	"reflect"
)

type NameEnum int

const (
	Alice NameEnum = iota + 1
	Bob
	Carol
	Dave
	Eve
)

func (name NameEnum) String() string {
	switch name {
	case Alice:
		return "Alice"
	case Bob:
		return "Bob"
	case Carol:
		return "Carol"
	case Dave:
		return "Dave"
	case Eve:
		return "Eve"
	default:
		return "unknown"
	}
}

type Animal struct {
}

func (a *Animal) Eat() {
	fmt.Println("Eat")
}

func doCall() {
	a := Animal{}
	reflect.ValueOf(&a).MethodByName("Eat").Call([]reflect.Value{})
}
